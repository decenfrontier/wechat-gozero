package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wslynn/wechat-gozero/common/biz"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		TransInsertSystemGroup(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn, c cache.CacheConf) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn, c),
	}
}

// 添加系统用户 组
func (m *defaultGroupModel) TransInsertSystemGroup(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error) {
	// 创建 与 微信团队 的群
	groupId1 := biz.GetGroupId(1, userId)
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, groupId1, "微信团队", GroupTypeSingle, GroupStatusYes, nil)
	if err != nil {
		return nil, err
	}
	// 创建 与 文件传输助手 的群
	groupId2 := biz.GetGroupId(2, userId)
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err = session.ExecCtx(ctx, query, groupId2, "文件传输助手", GroupTypeSingle, GroupStatusYes, nil)
	if err != nil {
		return nil, err
	}
	return ret, err
}
