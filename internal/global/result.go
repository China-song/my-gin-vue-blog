package global

import "fmt"

const (
	SUCCESS = 0
)

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
	OkResult = RegisterResult(SUCCESS, "OK")
)

var (
	ErrRequest = RegisterResult(9001, "请求参数格式错误")
	ErrDbOp    = RegisterResult(9004, "数据库操作异常")
	ErrRedisOp = RegisterResult(9005, "Redis 操作异常")

	ErrPassword     = RegisterResult(1002, "密码不正确")
	ErrUserNotExist = RegisterResult(1003, "该用户不存在")

	ErrTokenCreate = RegisterResult(1205, "TOKEN 生成失败")
	ErrPermission  = RegisterResult(1206, "权限不足")

	ErrTagHasArt  = RegisterResult(4003, "删除失败，标签下存在文章")
	ErrCateHasArt = RegisterResult(3003, "删除失败，分类下存在文章")
)
