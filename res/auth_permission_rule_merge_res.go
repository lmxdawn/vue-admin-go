package res

type AuthPermissionRuleMergeRes struct {
	Id        int64                        `json:"id"`        // ID
	Pid       int64                        `json:"pid"`       // 上级ID
	Name      string                       `json:"name"`      // 权限名
	Title     string                       `json:"title"`     // 权限标题
	Status    int                          `json:"status"`    // 状态（1：正常，0：禁用）
	Condition string                       `json:"condition"` // 规则表达式，为空表示存在就验证，不为空表示按照条件验证
	Listorder int64                        `json:"listorder"` // 排序，优先级，越小优先级越高
	Children  []AuthPermissionRuleMergeRes // 子节点（一次性加载所有节点时需要）
}
