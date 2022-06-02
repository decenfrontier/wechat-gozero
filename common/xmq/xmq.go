package xmq

import (
	"encoding/json"
	"ws_chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-queue/kq"
)

// 生产到消息队列
func PushToMq(producer *kq.Pusher, object interface{}) error {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MARSHAL_ERROR), "marshal object failed, object: %+v", object)
	}
	jsonStr := string(jsonBytes)
	// 放入消息队列
	err = producer.Push(jsonStr)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.MQ_ERROR), "push message to mq failed, jsonStr: %s", jsonStr)
	}
	return nil
}
