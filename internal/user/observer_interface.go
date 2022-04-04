package user

import "context"

type Observer interface {
	Observe(ctx context.Context) error
}
