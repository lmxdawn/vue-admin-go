package req

type AuthPermissionRuleSaveReq struct {
	Id        int64  `json:"id"`                                  // ID（编辑时必传）
	Pid       int64  `json:"pid" binding:"required"`              // 父级ID
	Name      string `json:"name" binding:"required"`             // 规则名称
	Title     string `json:"title" binding:"required"`            // 标题
	Status    int    `json:"status" binding:"required,oneof=0 1"` // 状态（1：正常，0：禁用）
	Condition string `json:"condition"`                           // 表达式
	Listorder int64  `json:"listorder"`                           // 排序
}
