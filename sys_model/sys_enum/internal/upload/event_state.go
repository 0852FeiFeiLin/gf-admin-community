package sys_enum_upload

import (
	"context"

	"github.com/kysion/base-library/utility/enum"
	"github.com/gogf/gf/v2/frame/g"
)

type EventStateEnum enum.IEnumCode[int]

type eventState struct {
	BeforeCache EventStateEnum
	AfterCache  EventStateEnum
	BeforeSave  EventStateEnum
	AfterSave   EventStateEnum
}

var EventState = eventState{
	BeforeCache: enum.New[EventStateEnum](1, "缓存前"),
	AfterCache:  enum.New[EventStateEnum](2, "已缓存"),
	BeforeSave:  enum.New[EventStateEnum](4, "保存前"),
	AfterSave:   enum.New[EventStateEnum](8, "保存后"),
}

func (e eventState) New(code int, description string) EventStateEnum {

	if (code&EventState.BeforeCache.Code()) == EventState.BeforeCache.Code() ||
		(code&EventState.AfterCache.Code()) == EventState.AfterCache.Code() ||
		(code&EventState.BeforeSave.Code()) == EventState.BeforeSave.Code() ||
		(code&EventState.AfterSave.Code()) == EventState.AfterSave.Code() {
		return enum.New[EventStateEnum](code, description)
	}
	// 记录无效的enum code，但返回一个默认值而不是panic
	g.Log().Error(context.Background(), "Upload.EventState.New: 无效的code", g.Map{"code": code, "description": description})
	return enum.New(code, description) // 返回原始的enum，让调用方决定如何处理
}
