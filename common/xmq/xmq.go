package xmq

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/pkg/errors"
)



// 生产到消息队列
func PushToMq(ctx context.Context, mqWriter *kafka.Writer, object interface{}) error {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MARSHAL_ERROR), "marshal object failed, err: %s, object: %+v", err, object)
	}
	// 放入消息队列
	err = mqWriter.WriteMessages(ctx, kafka.Message{Value: jsonBytes})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MQ_ERROR), "push message to mq failed, err: %s", err)
	}
	return nil
}
