package kvrocks

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"time"

	"github.com/daffarg/distributed-cascading-cb/repository"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/redis/go-redis/v9"
)

type kvRocks struct {
	client *redis.Client // using redis as client for Apache KVRocks
}

func NewKVRocksRepository(host, port, password string, db int) (repository.Repository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	if err := redisotel.InstrumentTracing(client); err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &kvRocks{
		client: client,
	}, nil
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

func (k *kvRocks) AddMembersIntoSet(ctx context.Context, key string, members ...string) (int64, error) {
	return k.client.SAdd(ctx, key, members).Result()
}

func (k *kvRocks) IsMemberOfSet(ctx context.Context, key, value string) (bool, error) {
	isMember, err := k.client.SIsMember(ctx, key, value).Result()
	if err != nil {
		return false, err
	}

	return isMember, err
}

func (k *kvRocks) IsMembersOfSet(ctx context.Context, key string, value ...string) ([]bool, error) {
	isMembers, err := k.client.SMIsMember(ctx, key, value).Result()
	if err != nil {
		return nil, err
	}

	return isMembers, nil
}

func (k *kvRocks) GetMemberOfSet(ctx context.Context, key string) ([]string, error) {
	members, err := k.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return members, err
}

func (k *kvRocks) IsKeyExist(ctx context.Context, key string) (bool, error) {
	isExist, err := k.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return isExist == 1, err
}

func (k *kvRocks) Scan(ctx context.Context, pattern string, count int64) ([]string, error) {
	var cursor uint64
	keys := make([]string, 0)
	for {
		var err error
		tmpKeys, cursor, err := k.client.Scan(ctx, cursor, pattern, count).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, tmpKeys...)
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}
