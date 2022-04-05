package main

import (
	"context"
	"github.com/kagestonedragon/client/configs"
	"github.com/kagestonedragon/client/internal/user"
	u "github.com/kagestonedragon/server/pkg/user"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l := log.New(os.Stdout, "", log.LstdFlags)

	cfg := configs.NewConfig()
	if err := cfg.Read(); err != nil {
		log.Fatalf("read config: %s", err)
	}

	c, err := grpc.Dial(cfg.Service.Target, grpc.WithInsecure())
	if err != nil {
		l.Fatal(err)
	}

	s := u.NewGRPCClient(c, nil, nil)

	var observer user.Observer
	switch cfg.Type {
	case "file":
		observer = createFileObserver(cfg, s, l)
		break
	case "nats":
		observer = createNatsObserver(cfg, s, l)
		break
	default:
		l.Fatalf("Type %s not supported", cfg.Type)
	}

	if err = observer.Observe(ctx); err != nil {
		l.Fatal(err)
	}
}

func createNatsObserver(cfg *configs.Config, service u.Service, logger *log.Logger) user.Observer {
	nc, err := nats.Connect(cfg.Observer.Nats.Target)
	if err != nil {
		logger.Fatal(err)
	}

	return user.NewNatsObserver(
		nc,
		cfg.Observer.Chunk,
		cfg.Observer.Nats.Subject,
		time.Duration(cfg.Observer.Pause)*time.Second,
		service,
		logger,
	)
}

func createFileObserver(cfg *configs.Config, service u.Service, logger *log.Logger) user.Observer {
	return user.NewFileObserver(
		cfg.Observer.File.Path,
		cfg.Observer.Chunk,
		time.Duration(cfg.Observer.Pause)*time.Second,
		service,
		logger,
	)
}
