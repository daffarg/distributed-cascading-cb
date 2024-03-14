package kvrocks

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/daffarg/distributed-cascading-cb/repository"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/redis/go-redis/v9"
)

type kvRocks struct {
	client *redis.Client // using redis as client for Apache KVRocks
}

func NewKVRocksRepository(host, port, password string, db int) repository.Repository {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	return &kvRocks{
		client: client,
	}
}

func (k *kvRocks) Set(ctx context.Context, key, value string) error {
	return k.client.Set(ctx, key, value, 0).Err()
}

func (k *kvRocks) SetWithExp(ctx context.Context, key, value string, exp time.Duration) error {
	return k.client.Set(ctx, key, value, exp).Err()
}

func (k *kvRocks) Get(ctx context.Context, key string) (string, error) {
	value, err := k.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", util.ErrKeyNotFound
		}
		return "", err
	}

	return value, err
}

func (k *kvRocks) AddMembersIntoSet(ctx context.Context, key string, members ...string) error {
	return k.client.SAdd(ctx, key, members).Err()
}

func (k *kvRocks) IsMemberOfSet(ctx context.Context, key, value string) (bool, error) {
	isMember, err := k.client.SIsMember(ctx, key, value).Result()
	if err != nil {
		return false, err
	}

	return isMember, err
}

func (k *kvRocks) GetMemberOfSet(ctx context.Context, key string) ([]string, error) {
	members, err := k.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return members, err
}
