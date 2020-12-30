package errcode

import "fmt"

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}
var codes = map[int]string{}

func NewError(code int,msg string) *Error{
	if _,ok := codes[code];ok{
		panic(fmt.Sprintf("错误码 %d已经存在,请更换一个",code))
	}
	codes[code] = msg
	return &Error{code:code,msg:msg}
}