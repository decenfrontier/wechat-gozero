package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// 群组类型
const GroupTypeSingle int64 = 1 // 单聊
const GroupTypeMulti int64 = 2  // 群聊

// 群组状态
const GroupStatusYes int64 = 1   // 正常
const GroupStatusNo int64 = 2    // 未通过
const GroupStatusBlack int64 = 3 // 黑名单
