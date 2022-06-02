package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/ds"
)

const servicePrefix = "bmstu_fv_parser."

func getUserKey(userVkID int) string {
	return servicePrefix + strconv.Itoa(userVkID)
}

type Client struct {
	cfg    config.RedisConfig
	client *redis.Client
}

func New(ctx context.Context) (*Client, error) {
	client := &Client{}
	cfg := config.FromContext(ctx)

	client.cfg = cfg.Redis

	redisClient := redis.NewClient(&redis.Options{
		Password:    cfg.Redis.Password,
		Username:    cfg.Redis.User,
		Addr:        cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		DB:          0,
		DialTimeout: time.Duration(cfg.Redis.DialTimeout) * time.Millisecond,
		ReadTimeout: time.Duration(cfg.Redis.ReadTimeout) * time.Millisecond,
	})

	client.client = redisClient

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("cant ping redis: %w", err)
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) SetUser(ctx context.Context, user ds.User) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(user); err != nil {
		return err
	}

	return c.client.Set(ctx, getUserKey(user.VkID), b.String(), 0).Err()
}

func (c *Client) GetUser(ctx context.Context, vkID int) (*ds.User, error) {
	data, err := c.client.Get(ctx, getUserKey(vkID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	b := strings.NewReader(data)

	res := &ds.User{}

	if err = json.NewDecoder(b).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
