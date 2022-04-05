package user

import (
	"context"
	"encoding/csv"
	"github.com/kagestonedragon/server/pkg/user"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type fileObserver struct {
	file    string
	chunk   int
	pause   time.Duration
	service user.Service
	logger  *log.Logger
}

func NewFileObserver(f string, c int, p time.Duration, s user.Service, l *log.Logger) Observer {
	return &fileObserver{
		file:    f,
		chunk:   c,
		pause:   p,
		service: s,
		logger:  l,
	}
}

func (o *fileObserver) Observe(ctx context.Context) error {
	t := time.Time{}

	for {
		s, err := os.Stat(o.file)
		if errors.Is(err, os.ErrNotExist) {
			time.Sleep(o.pause)
			continue
		} else if err != nil {
			return errors.Wrap(err, "read file stats")
		}

		if m := s.ModTime(); m.After(t) {
			t = m
		} else {
			time.Sleep(o.pause)
			continue
		}

		f, err := os.Open(o.file)
		if err != nil {
			return errors.Wrap(err, "open file")
		}

		o.process(ctx, csv.NewReader(f))

		if err := f.Close(); err != nil {
			return errors.Wrap(err, "close file")
		}
	}
}

func (o *fileObserver) process(ctx context.Context, r *csv.Reader) {
	var wg sync.WaitGroup
	eof := false

	for {
		for i := 1; i <= o.chunk; i++ {
			record, err := r.Read()
			if record == nil {
				eof = true
				break
			}
			if err != nil {
				o.logger.Printf("Read error: %s\n", err)
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				o.add(ctx, record)
			}()
		}
		wg.Wait()

		if eof == true {
			break
		}
	}
}

func (o *fileObserver) parse(u []string) (*user.AddUserRequest, error) {
	if len(u) != 2 || len(u[0]) <= 0 {
		return nil, errors.New("wrong format")
	}

	a, err := strconv.ParseBool(u[1])
	if err != nil {
		return nil, errors.New("cannot get active value")
	}

	return &user.AddUserRequest{
		Name:   u[0],
		Active: a,
	}, nil
}

func (o *fileObserver) add(ctx context.Context, u []string) {
	r, err := o.parse(u)
	if err != nil {
		o.logger.Printf("Parse error: %s\n", err)
		return
	}

	if u, err := o.service.AddUser(ctx, r); err != nil {
		o.logger.Printf("Service error: %s\n", err)
	} else {
		o.logger.Printf("User add: %s - %d\n", u.Name, u.Id)
	}
}
