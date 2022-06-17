package xmq

import (
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/pkg/errors"
)

// 生产到消息队列
func PushToMq(conn *kafka.Conn, object interface{}) error {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MARSHAL_ERROR), "marshal object failed, object: %+v", object)
	}
	// 放入消息队列
	_, err = conn.WriteMessages(kafka.Message{Value: jsonBytes})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MQ_ERROR), "push message to mq failed, object: %+v", object)
	}
	return nil
}
