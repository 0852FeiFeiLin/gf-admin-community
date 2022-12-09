package kyUser

import "github.com/SupenBysz/gf-admin-community/model/enum"

type TypeEnum kyEnum.Code

type userType struct {
	Anonymous   TypeEnum
	User        TypeEnum
	WeBusiness  TypeEnum
	Merchant    TypeEnum
	Advertiser  TypeEnum
	Facilitator TypeEnum
	Operator    TypeEnum
	SuperAdmin  TypeEnum
}

var Type = userType{
	Anonymous:   kyEnum.New(0, "匿名"),
	User:        kyEnum.New(1, "用户"),
	WeBusiness:  kyEnum.New(2, "微商"),
	Merchant:    kyEnum.New(4, "商户"),
	Advertiser:  kyEnum.New(8, "广告主"),
	Facilitator: kyEnum.New(16, "服务商"),
	Operator:    kyEnum.New(32, "运营商"),
	SuperAdmin:  kyEnum.New(-1, "超级管理员"),
}

func (e userType) New(code int, description string) TypeEnum {
	if (code&Type.Anonymous.Code()) == Type.Anonymous.Code() ||
		(code&Type.User.Code()) == Type.User.Code() ||
		(code&Type.WeBusiness.Code()) == Type.WeBusiness.Code() ||
		(code&Type.Merchant.Code()) == Type.Merchant.Code() ||
		(code&Type.Advertiser.Code()) == Type.Advertiser.Code() ||
		(code&Type.Facilitator.Code()) == Type.Facilitator.Code() ||
		(code&Type.Operator.Code()) == Type.Operator.Code() ||
		(code&Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
		return kyEnum.New(code, description)
	}
	panic("UserType: error")
}
