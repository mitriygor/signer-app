package listener

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	IncrCount()
	SetCount(count int)
	GetCount() int
}

type ListenerRepository struct {
	redisClient *redis.Client
}

func NewListenerRepository(redisClient *redis.Client) Repository {
	return &ListenerRepository{
		redisClient: redisClient,
	}
}

func (lr *ListenerRepository) IncrCount() {
	lr.redisClient.Incr(context.TODO(), "listener_count")
}

func (lr *ListenerRepository) SetCount(count int) {
	lr.redisClient.Set(context.TODO(), "listener_count", count, 0)
}

func (lr *ListenerRepository) GetCount() int {
	count, err := lr.redisClient.Get(context.TODO(), "listener_count").Int()
	if err != nil {
		fmt.Printf("\nERedis :: GetCount :: ERROR:%v\n", err.Error())
		return -1
	}
	return count
}
