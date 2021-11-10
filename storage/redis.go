package storage

import (
	"context"
	"errors"
	"strconv"
	"time"
	"urlshortener/utils"

	"github.com/go-redis/redis/v8"
)

type redisdb struct{ client *redis.Client }

var ctx = context.Background()

func New(host string, port string, db int) (Service, error) {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		DB:   db,
	})
	return &redisdb{client}, nil
}

func (r *redisdb) isUsed(id string) bool {
	exists, err := r.client.Exists(ctx, id).Result()
	if err != nil {
		return false
	}
	return exists > 0
}

func (r *redisdb) getUniqueId() (string, error) {
	baseKey, err := r.client.Incr(ctx, utils.BaseKeyName).Result()
	if err != nil {
		return "", err
	}
	if baseKey <= 0 {
		return "", errors.New("BaseKey has non-positive value")
	}
	var uniqueId string
	for used := true; used; used = r.isUsed(uniqueId) {
		uniqueId = utils.GenerateUniqueId(uint64(baseKey))
	}
	return uniqueId, nil
}

func (r *redisdb) Add(url string) (*Item, error) {
	id, err := r.getUniqueId()
	if err != nil {
		return nil, err
	}
	createdAt := time.Now().Unix()
	_, err = r.client.HMSet(ctx, id, map[string]interface{}{
		"url":       url,
		"createdat": createdAt,
		"visits":    0,
	}).Result()
	if err != nil {
		return nil, err
	}
	return &Item{id, url, createdAt, 0}, nil
}

func (r *redisdb) Visit(id string) (string, error) {
	url, err := r.client.HGet(ctx, id, "url").Result()
	if err != nil {
		return "/", err
	}
	err = r.client.HIncrBy(ctx, id, "visits", 1).Err()
	if err != nil {
		return "/", err
	}
	if len(url) > 0 {
		return url, nil
	}
	return "/", nil
}

func (r *redisdb) GetInfo(id string) (*Item, error) {
	itemMap, err := r.client.HGetAll(ctx, id).Result()
	if err != nil {
		return nil, err
	}
	if len(itemMap["url"]) > 0 {
		createdAt, err := strconv.ParseInt(itemMap["createdat"], 10, 64)
		if err != nil {
			return nil, err
		}
		visits, err := strconv.ParseInt(itemMap["visits"], 10, 64)
		if err != nil {
			return nil, err
		}
		return &Item{id, itemMap["url"], createdAt, visits}, nil
	}
	return nil, nil
}

func (r *redisdb) Close() error {
	return r.client.Close()
}
