package database

import (
	"back-end/config"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis(config *config.Config) *Redis {
	return &Redis{
		rdb: ConnectRedis(config),
	}
}

func ConnectRedis(config *config.Config) *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       0,
	})
}

func (r Redis) Get(key string) (string, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r Redis) Set(key string, entry string) error {
	return r.rdb.Set(context.Background(), key, entry, 15*time.Minute).Err()
}

func (r Redis) Del(key string) error {
	_, err := r.rdb.Del(context.Background(), key).Result()
	return err
}

func (r Redis) RPush(key string, values ...interface{}) error {
	_, err := r.rdb.RPush(context.Background(), key, values...).Result()
	return err
}

func (r Redis) LPop(key string) ([]byte, error) {
	val, err := r.rdb.LPop(context.Background(), key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r Redis) LIndex(key string, index int64) ([]byte, error) {
	val, err := r.rdb.LIndex(context.Background(), key, index).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r Redis) LRange(key string, start, stop int64) ([][]byte, error) {
	vals, err := r.rdb.LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	byteVals := make([][]byte, len(vals))
	for i, val := range vals {
		byteVals[i] = []byte(val)
	}
	return byteVals, nil
}

func (r Redis) LSet(key string, index int64, value string) error {
	err := r.rdb.LSet(context.Background(), key, index, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) LRem(key string, count int64, value string) error {
	err := r.rdb.LRem(context.Background(), key, count, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) AddToSet(key string, members ...string) error {
	return r.rdb.SAdd(context.Background(), key, members).Err()
}

func (r Redis) GetSetMembers(key string) ([]string, error) {
	return r.rdb.SMembers(context.Background(), key).Result()
}

func (r Redis) RemoveFromSet(key string, members ...string) error {
	return r.rdb.SRem(context.Background(), key, members).Err()
}
