package main

import (
	"context"
	"github.com/kagestonedragon/client/configs"
	"github.com/kagestonedragon/client/internal/user"
	u "github.com/kagestonedragon/server/pkg/user"
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

	observer := user.NewFileObserver(
		cfg.Observer.File,
		cfg.Observer.Chunk,
		time.Duration(cfg.Observer.Pause)*time.Second,
		u.NewGRPCClient(c, nil, nil),
		l,
	)

	if err = observer.Observe(ctx); err != nil {
		l.Fatal(err)
	}
}
