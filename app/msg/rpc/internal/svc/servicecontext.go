package svc

import (
	"github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/config"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ChatMsgModel model.ChatMsgModel
	RedisClient  *redis.Redis
	MqProducer   *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config:       c,
		ChatMsgModel: model.NewChatMsgModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		MqProducer: kq.NewPusher(c.MqConf.Brokers, c.MqConf.Topic),
	}
}
