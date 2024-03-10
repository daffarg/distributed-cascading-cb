package repository

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/repository"
	"github.com/ledisdb/ledisdb/config"
	"github.com/ledisdb/ledisdb/ledis"
)

type ledisDB struct {
	db *ledis.DB
}

func NewLedisDB(idx int) repository.Repository {
	cfg := config.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(idx)

	return &ledisDB{db: db}
}

func (l *ledisDB) Set(ctx context.Context, key, value string) error {
	return l.db.Set([]byte(key), []byte(value))
}

func (l *ledisDB) Get(ctx context.Context, key string) (string, error) {
	value, err := l.db.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(value[:]), nil
}
