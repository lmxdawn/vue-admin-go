package model

import "time"

type AuthRole struct {
	Id         int        //
	Name       string     // 角色名称
	Pid        int        // 父角色ID
	Status     int8       // 状态
	Remark     string     // 备注
	CreateTime *time.Time // 创建时间
	UpdateTime *time.Time // 更新时间
	Listorder  int        // 排序，优先级，越小优先级越高
}
