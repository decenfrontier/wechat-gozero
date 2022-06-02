package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wslynn/wechat-gozero/common/biz"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupUserModel = (*customGroupUserModel)(nil)

type (
	// GroupUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupUserModel.
	GroupUserModel interface {
		groupUserModel
		FindUserListByGroupId(ctx context.Context, groupId string) ([]*GroupUser, error)
		FindUserIdListByGroupId(ctx context.Context, groupId string) ([]int64, error)
		FindGroupIdListByUserId(ctx context.Context, userId int64) ([]string, error)
		FindAliasNameByGroupAndUser(ctx context.Context, groupId string, userId int64) (string, error)
		TransInsertSystemGroupUser(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error)
	}

	customGroupUserModel struct {
		*defaultGroupUserModel
	}
)

// NewGroupUserModel returns a model for the database table.
func NewGroupUserModel(conn sqlx.SqlConn, c cache.CacheConf) GroupUserModel {
	return &customGroupUserModel{
		defaultGroupUserModel: newGroupUserModel(conn, c),
	}
}

func (m *defaultGroupUserModel) FindUserListByGroupId(ctx context.Context, groupId string) ([]*GroupUser, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ?", groupUserRows, m.table)
	var resp []*GroupUser
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultGroupUserModel) FindUserIdListByGroupId(ctx context.Context, groupId string) ([]int64, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ?", "user_id", m.table)
	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultGroupUserModel) FindGroupIdListByUserId(ctx context.Context, userId int64) ([]string, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", "group_id", m.table)
	var resp []string
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 添加系统用户 组用户
func (m *defaultGroupUserModel) TransInsertSystemGroupUser(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error) {
	// 添加 微信团队 为好友
	groupId1 := biz.GetGroupId(1, userId)
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, groupUserRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, groupId1, userId, "微信团队")
	if err != nil {
		return nil, err
	}
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, groupUserRowsExpectAutoSet)
	ret, err = session.ExecCtx(ctx, query, groupId1, 1, fmt.Sprintf("%d", userId))
	if err != nil {
		return nil, err
	}

	// 添加 文件传输助手 为好友
	groupId2 := biz.GetGroupId(2, userId)
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, groupUserRowsExpectAutoSet)
	ret, err = session.ExecCtx(ctx, query, groupId2, userId, "文件传输助手")
	if err != nil {
		return nil, err
	}
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, groupUserRowsExpectAutoSet)
	ret, err = session.ExecCtx(ctx, query, groupId2, 2, fmt.Sprintf("%d", userId))
	if err != nil {
		return nil, err
	}
	return ret, err
}

// 通过 groupId 和 userId 查找 用户对该群的备注
func (m *defaultGroupUserModel) FindAliasNameByGroupAndUser(ctx context.Context, groupId string, userId int64) (string, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ? and `user_id` = ?", "alias_name", m.table)
	var aliasName string
	err := m.QueryRowNoCacheCtx(ctx, &aliasName, query, groupId, userId)
	if err != nil {
		return "", err
	}
	return aliasName, nil
}
