package sys_enum_user

import (
	"context"

	"github.com/kysion/base-library/utility/enum"
	"github.com/gogf/gf/v2/frame/g"
)

type TypeEnum enum.IEnumCode[int]

type userType struct {
    Admin TypeEnum

    SuperAdmin TypeEnum
}

var Type = userType{
    Admin:      enum.New[TypeEnum](64, "后台"),
    SuperAdmin: enum.New[TypeEnum](-1, "超级管理员"),
}

func (e userType) New(code int, description string) TypeEnum {
    if (code & Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
        return Type.SuperAdmin
    }
    if (code & Type.Admin.Code()) == Type.Admin.Code() {
        return Type.Admin
    }
    // 记录无效的enum code，但返回一个默认值而不是panic
    g.Log().Error(context.Background(), "User.Type.New: 无效的code", g.Map{"code": code, "description": description})
    return enum.New[TypeEnum](code, description) // 返回原始的enum，让调用方决定如何处理
}
