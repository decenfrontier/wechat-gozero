package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatMsgModel = (*customChatMsgModel)(nil)

type (
	// ChatMsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatMsgModel.
	ChatMsgModel interface {
		chatMsgModel
		FindMsgListByLastMsgId(ctx context.Context, groupId string, minMsgId int64, maxMsgId int64) ([]*ChatMsg, error)
		FindLastMsgByGroupId(ctx context.Context, groupId string) (*ChatMsg, error)
	}

	customChatMsgModel struct {
		*defaultChatMsgModel
	}
)

// NewChatMsgModel returns a model for the database table.
func NewChatMsgModel(conn sqlx.SqlConn, c cache.CacheConf) ChatMsgModel {
	return &customChatMsgModel{
		defaultChatMsgModel: newChatMsgModel(conn, c),
	}
}

// 获取指定群组的离线消息列表
func (m *customChatMsgModel) FindMsgListByLastMsgId(ctx context.Context, groupId string, minMsgId int64, maxMsgId int64) ([]*ChatMsg, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ? and `id` > ? and `id` < ? order by `id` desc limit %d", chatMsgRows, m.table, PerPullNum)
	var resp []*ChatMsg
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId, minMsgId, maxMsgId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取指定群组的 最后一条消息
func (m *customChatMsgModel) FindLastMsgByGroupId(ctx context.Context, groupId string) (*ChatMsg, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ? order by `id` desc limit 1", chatMsgRows, m.table)
	var resp ChatMsg
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, groupId)
	if err != nil {
		return nil, err
	}
	return &resp, err
}