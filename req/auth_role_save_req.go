package req

type AuthRoleSaveReq struct {
	Id        int64  `json:"id" binding:"required"`               // 权限ID
	Pid       int64  `json:"pid" binding:"required"`              // 权限上级ID
	Name      string `json:"name" binding:"required"`             // 权限名
	Status    int    `json:"status" binding:"required,oneof=0 1"` // 状态（1：正常，0：禁用）
	Remark    string `json:"remark" binding:"required"`           // 备注
	Listorder int64  `json:"listorder" binding:"required"`        // 排序
}
