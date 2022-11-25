package model

type SysRole struct {
	Id          int64  `p:"id"         dc:"ID"`
	Name        string `p:"name"       v:"required|length:1,16#请输入角色名称|角色名称长度限定1~16字符" dc:"角色名称"`
	Description string `p:"description"       v:"max-length:128#角色描述长度限定128字符" dc:"角色描述"`
}
