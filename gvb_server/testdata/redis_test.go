package testdata

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "20.tcp.cpolar.top:13407",
		Password: "",
		DB:       db,
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis connect error:%v", err)
		return nil
	}
	return rdb
}
func TestRedis(t *testing.T) {
	rdb := ConnectRedis()
	res := rdb.Keys("*")
	s, _ := res.Result()
	t.Log(s)
}
