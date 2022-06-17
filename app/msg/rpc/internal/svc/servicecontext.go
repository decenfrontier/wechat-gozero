package svc

import (
	"context"

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
	MqConn       *kafka.Conn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	mqConn, err := kafka.DialLeader(context.TODO(), "tcp", c.MqConf.Brokers[0], c.MqConf.Topic, 0)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		ChatMsgModel: model.NewChatMsgModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		MqConn: mqConn,
	}
}
