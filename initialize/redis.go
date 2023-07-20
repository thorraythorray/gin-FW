package initialize

import (
	"github.com/thorraythorray/go-proj/global"

	"github.com/go-redis/redis/v8"
)

func RedisInit() {
	r := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host,
		Password: r.Passwd, // no password set
		DB:       r.DB,     // use default DB
	})
	global.Redis = client
}
