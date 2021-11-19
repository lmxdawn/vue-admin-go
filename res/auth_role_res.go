package res

type AuthRoleRes struct {
	Id        int64  `json:"id"`        // 角色ID
	Name      string `json:"name"`      // 角色名称
	Pid       int64  `json:"pid"`       // 角色父级ID
	Status    int    `json:"status"`    // 状态（1：正常，0：禁用）
	Remark    string `json:"remark"`    // 备注
	Listorder int64  `json:"listorder"` // 排序，优先级，越小优先级越高
}
