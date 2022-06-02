package xerr

import (
	"fmt"
)

type CodeError struct {
	code uint32
	msg  string
}

func (e *CodeError) GetCode() uint32 {
	return e.code
}

func (e *CodeError) GetMsg() string {
	return e.msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

func NewErrCodeMsg(code uint32, msg string) error {
	return &CodeError{code: code, msg: msg}
}

func NewErrCode(code uint32) error {
	return &CodeError{code: code, msg: mapCodeMsg[code]}
}

func NewErrMsg(msg string) error {
	return &CodeError{code: SERVER_ERROR, msg: msg}
}