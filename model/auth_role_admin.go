package model

type AuthRoleAdmin struct {
	RoleId  int // 角色 id
	AdminId int // 管理员id
}

func AuthRoleAdminListByAdminId(adminId int64) ([]AuthRoleAdmin, error) {
	var authRoleAdmins []AuthRoleAdmin
	db := DB.Select("role_id")
	err := db.Where("admin_id = ?", adminId).Find(&authRoleAdmins).Error
	if err != nil {
		return nil, err
	}
	return authRoleAdmins, nil
}

func AuthRoleAdminListByAdminIdIn(adminIds []int64) ([]AuthRoleAdmin, error) {
	var authRoleAdmins []AuthRoleAdmin
	db := DB.Select("role_id,admin_id")
	err := db.Where("admin_id IN ?", adminIds).Find(&authRoleAdmins).Error
	if err != nil {
		return nil, err
	}
	return authRoleAdmins, nil
}

func AuthRoleAdminListByRoleId(roleId int64) ([]AuthRoleAdmin, error) {
	var authRoleAdmins []AuthRoleAdmin
	db := DB.Select("admin_id")
	err := db.Where("role_id = ?", roleId).Find(&authRoleAdmins).Error
	if err != nil {
		return nil, err
	}
	return authRoleAdmins, nil
}

func AuthRoleAdminInsertAll(authRoleAdmins []AuthRoleAdmin) error {
	var data []map[string]interface{}
	for _, value := range authRoleAdmins {
		_ = append(data, map[string]interface{}{
			"role_id":  value.RoleId,
			"admin_id": value.AdminId,
		})
	}
	err := DB.Model(&AuthAdmin{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthRoleAdminDeleteById(id int64) error {
	err := DB.Where("id = ?", id).Delete(&AuthAdmin{}).Error
	if err != nil {
		return err
	}
	return nil
}
