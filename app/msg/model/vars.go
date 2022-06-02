package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const PerPullNum int64 = 10  // 每次消息拉取的数量

const MsgTypeText int64 = 1  // 文本
const MsgTypePic int64 = 2  // 图片
const MsgTypeVideo int64 = 3  // 视频
const MsgTypeAudio int64 = 4  // 音频

const MsgSayHello string = "我们已经是好友了, 现在可以开始聊天了"