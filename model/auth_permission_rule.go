package model

import "time"

type AuthPermissionRule struct {
	Id         int        // 规则编号
	Pid        int        // 父级id
	Name       string     // 规则唯一标识
	Title      string     // 规则中文名称
	Status     int8       // 状态：为1正常，为0禁用
	Condition  string     // 规则表达式，为空表示存在就验证，不为空表示按照条件验证
	Listorder  int        // 排序，优先级，越小优先级越高
	CreateTime *time.Time // 创建时间
	UpdateTime *time.Time // 更新时间
}
