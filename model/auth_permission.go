package model

type AuthPermission struct {
	RoleId           int    // 角色
	PermissionRuleId int    // 权限id
	Type             string // 权限规则分类，请加应用前缀,如admin_
}

func AuthPermissionListByRoleIdIn(roleIds []int64) ([]AuthPermission, error) {
	var authPermissions []AuthPermission
	db := DB.Select("permission_rule_id")
	err := db.Where("role_id IN ?", roleIds).Find(&authPermissions).Error
	if err != nil {
		return nil, err
	}
	return authPermissions, nil
}

func AuthPermissionListByRoleId(roleId int64) ([]AuthPermission, error) {
	var authPermissions []AuthPermission
	db := DB.Select("permission_rule_id")
	err := db.Where("role_id = ?", roleId).Find(&authPermissions).Error
	if err != nil {
		return nil, err
	}
	return authPermissions, nil
}

func AuthPermissionInsertAll(authPermissions []AuthPermission) error {
	var data []map[string]interface{}
	for _, value := range authPermissions {
		_ = append(data, map[string]interface{}{
			"role_id":            value.RoleId,
			"permission_rule_id": value.PermissionRuleId,
			"type":               value.Type,
		})
	}
	err := DB.Model(&AuthAdmin{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthPermissionDeleteByRoleId(roleId int64) error {
	err := DB.Where("role_id = ?", roleId).Delete(&AuthPermission{}).Error
	if err != nil {
		return err
	}
	return nil
}
