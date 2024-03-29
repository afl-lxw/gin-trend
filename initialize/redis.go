package initialize

import (
	"context"

	"github.com/afl-lxw/gin-trend/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.TREND_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.TREND_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.TREND_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.TREND_REDIS = client
	}
}
