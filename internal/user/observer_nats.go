package user

import (
	"context"
	"encoding/json"
	"github.com/kagestonedragon/server/pkg/user"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

type natsObserver struct {
	nc      *nats.Conn
	chunk   int
	pause   time.Duration
	service user.Service
	logger  *log.Logger
}

func NewNatsObserver(nc *nats.Conn, c int, p time.Duration, s user.Service, l *log.Logger) Observer {
	return &natsObserver{
		nc:      nc,
		chunk:   c,
		pause:   p,
		service: s,
		logger:  l,
	}
}

func (o *natsObserver) Observe(ctx context.Context) error {
	var wg sync.WaitGroup

	sub, err := o.nc.SubscribeSync("users")
	if err != nil {
		o.logger.Fatal(err)
	}
	o.logger.Printf("Open subscriber\n")

	for {
		for i := 1; i <= o.chunk; i++ {
			msg, err := sub.NextMsg(time.Duration(100) * time.Second)
			if err != nil {
				o.logger.Printf("Next message: %s\n", err)
				break
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				defer msg.Ack()

				o.process(ctx, msg)
			}()
		}
		wg.Wait()
	}

	return sub.Drain()
}

func (o *natsObserver) process(ctx context.Context, msg *nats.Msg) {
	var r user.AddUserRequest

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		o.logger.Printf("Unmarshal error: %s\n", err)
		return
	}

	if u, err := o.service.AddUser(ctx, &r); err != nil {
		o.logger.Printf("Service error: %s\n", err)
	} else {
		o.logger.Printf("User add: %s - %d\n", u.Name, u.Id)
	}
}
