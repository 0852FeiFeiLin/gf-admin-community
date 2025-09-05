package sys_enum_auth

import (
	"context"

	"github.com/kysion/base-library/utility/enum"
	"github.com/gogf/gf/v2/frame/g"
)

type ActionTypeEnum enum.IEnumCode[int]

type actionType struct {
	Login    ActionTypeEnum
	Logout   ActionTypeEnum
	Register ActionTypeEnum
}

var ActionType = actionType{
	Login:    enum.New[ActionTypeEnum](1, "登录"),
	Logout:   enum.New[ActionTypeEnum](2, "退出"),
	Register: enum.New[ActionTypeEnum](4, "注册"),
}

func (e actionType) New(code int, description string) ActionTypeEnum {
	if (code&ActionType.Login.Code()) == ActionType.Login.Code() ||
		(code&ActionType.Logout.Code()) == ActionType.Logout.Code() ||
		(code&ActionType.Register.Code()) == ActionType.Register.Code() {
		return enum.New[ActionTypeEnum](code, description)
	}
	// 记录无效的enum code，但返回一个默认值而不是panic
	g.Log().Error(context.Background(), "Auth.ActionType.New: 无效的code", g.Map{"code": code, "description": description})
	return enum.New(code, description) // 返回原始的enum，让调用方决定如何处理
}
