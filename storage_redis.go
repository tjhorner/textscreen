package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorageAdapter struct {
	redis         *redis.Client
	Context       context.Context
	SessionExpiry time.Duration
}

func (c *RedisStorageAdapter) hashKey(phoneNumber string) string {
	h := sha256.New()
	h.Write([]byte(phoneNumber))
	return string(h.Sum(nil))
}

func (c *RedisStorageAdapter) Delete(conversation *Conversation) error {
	key := c.hashKey(conversation.PhoneNumber)
	return c.redis.Del(c.Context, key).Err()
}

func (c *RedisStorageAdapter) Save(conversation *Conversation) error {
	key := c.hashKey(conversation.PhoneNumber)

	encoded, err := json.Marshal(conversation)
	if err != nil {
		return err
	}

	err = c.redis.Set(c.Context, key, string(encoded), c.SessionExpiry).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *RedisStorageAdapter) Get(phoneNumber string) (*Conversation, error) {
	key := c.hashKey(phoneNumber)

	val, err := c.redis.Get(c.Context, key).Result()
	if err != nil {
		return nil, err
	}

	var result Conversation
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *RedisStorageAdapter) Exists(phoneNumber string) (bool, error) {
	key := c.hashKey(phoneNumber)

	exists, err := c.redis.Exists(c.Context, key).Result()
	if err != nil {
		return false, err
	}

	if exists == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func NewRedisStorageAdapter(redis *redis.Client) *RedisStorageAdapter {
	return &RedisStorageAdapter{
		redis:         redis,
		Context:       context.Background(),
		SessionExpiry: 24 * time.Hour,
	}
}
