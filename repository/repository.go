package repository

import (
	"context"
	"time"
)

type Repository interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	SetWithExp(ctx context.Context, key, value string, exp time.Duration) error
	AddMembersIntoSet(ctx context.Context, key string, members ...string) (int64, error)
	IsMemberOfSet(ctx context.Context, key, value string) (bool, error)
	IsMembersOfSet(ctx context.Context, key string, value ...string) ([]bool, error)
	GetMemberOfSet(ctx context.Context, key string) ([]string, error)
}
