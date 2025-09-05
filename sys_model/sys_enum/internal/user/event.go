package sys_enum_user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/enum"
)

type EventEnum enum.IEnumCode[int]

type event struct {
	BeforeCreate   EventEnum
	AfterCreate    EventEnum
	ChangePassword EventEnum
	ResetPassword  EventEnum
	ResetEmail     EventEnum
}

var Event = event{
	BeforeCreate:   enum.New[EventEnum](2, "创建前"),
	AfterCreate:    enum.New[EventEnum](4, "创建后"),
	ChangePassword: enum.New[EventEnum](8, "修改密码"),
	ResetPassword:  enum.New[EventEnum](16, "重置密码"),
	ResetEmail:     enum.New[EventEnum](16, "重置邮箱"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() ||
		(code&Event.ChangePassword.Code()) == Event.ChangePassword.Code() ||
		(code&Event.ResetPassword.Code()) == Event.ResetPassword.Code() ||
		(code&Event.ResetEmail.Code()) == Event.ResetEmail.Code() {
		return enum.New[EventEnum](code, description)
	}
	// 记录无效的enum code，但返回一个默认值而不是panic
	g.Log().Error(context.Background(), "User.Event.New: 无效的code", g.Map{"code": code, "description": description})
	return enum.New[EventEnum](code, description) // 返回原始的enum，让调用方决定如何处理
}
