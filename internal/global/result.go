package global

import "fmt"

type Result struct {
	code int
	msg  string
}

func (r Result) Code() int {
	return r.code
}

func (r Result) Msg() string {
	return r.msg
}

var (
	//_codes    = map[int]struct{}{}
	_messages = make(map[int]string) // 错误码: 错误消息
)

func RegisterResult(code int, msg string) Result {
	if _, ok := _messages[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	if msg == "" {
		panic("错误码消息不能为空")
	}

	_messages[code] = msg

	return Result{
		code: code,
		msg:  msg,
	}
}

var (
	ErrRequest = RegisterResult(9001, "请求参数格式错误")
	ErrDbOp    = RegisterResult(9004, "数据库操作异常")

	ErrPassword     = RegisterResult(1002, "密码不正确")
	ErrUserNotExist = RegisterResult(1003, "该用户不存在")
)
