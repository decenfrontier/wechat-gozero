package svc

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ChatMsgModel model.ChatMsgModel
	RedisClient  *redis.Redis
	MqWriter     *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	mqWriter := &kafka.Writer{
		Addr:         kafka.TCP(c.MqConf.Brokers...),
		Topic:        c.MqConf.Topic,
		BatchTimeout: time.Millisecond * 20,
	}
	return &ServiceContext{
		Config:       c,
		ChatMsgModel: model.NewChatMsgModel(sqlConn, c.Cache),
		RedisClient:  redisClient,
		MqWriter:     mqWriter,
	}
}
