package sys_enum_auth

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type ActionTypeEnum enum.Code

type actionType struct {
	Login    ActionTypeEnum
	Logout   ActionTypeEnum
	Register ActionTypeEnum
}

var ActionType = actionType{
	Login:    enum.New(1, "登录"),
	Logout:   enum.New(2, "退出"),
	Register: enum.New(4, "注册"),
}

func (e actionType) New(code int, description string) ActionTypeEnum {
	if (code&ActionType.Login.Code()) == ActionType.Login.Code() ||
		(code&ActionType.Logout.Code()) == ActionType.Logout.Code() ||
		(code&ActionType.Register.Code()) == ActionType.Register.Code() {
		return enum.New(code, description)
	}
	panic("Auth.ActionType.New: error")
}
