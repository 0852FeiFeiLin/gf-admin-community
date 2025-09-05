package sys_enum_user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/enum"
)

type StateEnum enum.IEnumCode[int]

type state struct {
	Unactivated StateEnum
	Normal      StateEnum
	Suspended   StateEnum
	Abnormality StateEnum
	Canceled    StateEnum
}

var State = state{
	Unactivated: enum.New[StateEnum](0, "未激活"),
	Normal:      enum.New[StateEnum](1, "正常"),
	Suspended:   enum.New[StateEnum](-1, "封号"),
	Abnormality: enum.New[StateEnum](-2, "异常"),
	Canceled:    enum.New[StateEnum](-3, "已注销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Unactivated.Code()) == State.Unactivated.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.Suspended.Code()) == State.Suspended.Code() ||
		(code&State.Abnormality.Code()) == State.Abnormality.Code() ||
		(code&State.Canceled.Code()) == State.Canceled.Code() {
		return enum.New[StateEnum](code, description)
	}
	// 记录无效的enum code，但返回一个默认值而不是panic
	g.Log().Error(context.Background(), "User.State.New: 无效的code", g.Map{"code": code, "description": description})
	return enum.New[StateEnum](code, description) // 返回原始的enum，让调用方决定如何处理
}
