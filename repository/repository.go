package repository

import (
	"context"
)

type Repository interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}
