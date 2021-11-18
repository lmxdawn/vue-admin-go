package model

type AuthPermission struct {
	RoleId           int    // 角色
	PermissionRuleId int    // 权限id
	Type             string // 权限规则分类，请加应用前缀,如admin_
}
