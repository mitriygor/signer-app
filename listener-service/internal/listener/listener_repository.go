package listener

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	IncrCount(countName string)
	SetCount(count int, countName string)
	GetCount(countName string) int
}

type ListenerRepository struct {
	redisClient *redis.Client
}

func NewListenerRepository(redisClient *redis.Client) Repository {
	return &ListenerRepository{
		redisClient: redisClient,
	}
}

func (lr *ListenerRepository) IncrCount(countName string) {
	lr.redisClient.Incr(context.TODO(), countName)
}

func (lr *ListenerRepository) SetCount(count int, countName string) {
	lr.redisClient.Set(context.TODO(), countName, count, 0)
}

func (lr *ListenerRepository) GetCount(countName string) int {
	count, err := lr.redisClient.Get(context.TODO(), countName).Int()
	if err != nil {
		fmt.Printf("\nListener :: Redis :: GetCount :: error:%v\n", err.Error())
		return -1
	}
	return count
}
