package model

import (
	"fmt"
	"github.com/lmxdawn/vue-admin-go/req"
	"time"
)

type AuthRole struct {
	Id         int        //
	Name       string     // 角色名称
	Pid        int        // 父角色ID
	Status     int        // 状态
	Remark     string     // 备注
	CreateTime *time.Time // 创建时间
	UpdateTime *time.Time // 更新时间
	Listorder  int        // 排序，优先级，越小优先级越高
}

func AuthRoleListPage(q req.AuthRoleQueryReq) ([]AuthRole, error) {
	var authRoles []AuthRole
	offset := (q.Page - 1) * q.Limit
	db := DB.Select("id,`name`,status,remark,create_time,listorder")
	if q.Status != -1 {
		db = db.Where("status = ? ", q.Status)
	}
	if q.Name != "" {
		db = db.Where("name LIKE ? ", fmt.Sprint("%", q.Name, "%"))
	}
	err := db.Offset(offset).Limit(q.Limit).Order("listorder DESC").Find(&authRoles).Error
	if err != nil {
		return nil, err
	}
	return authRoles, nil
}

func AuthRoleListRolePage(q req.AuthRoleQueryReq) ([]AuthRole, error) {
	var authRoles []AuthRole
	offset := (q.Page - 1) * q.Limit
	db := DB.Select("id,`name`")
	if q.Status != -1 {
		db = db.Where("status = ? ", q.Status)
	}
	err := db.Offset(offset).Limit(q.Limit).Order("listorder DESC").Find(&authRoles).Error
	if err != nil {
		return nil, err
	}
	return authRoles, nil
}

func AuthRoleFindByName(name string) (*AuthRole, error) {
	var authRole *AuthRole
	db := DB.Select("id")
	err := db.Where("name = ?", name).First(&authRole).Error
	if err != nil {
		return nil, err
	}
	return authRole, nil
}

func AuthRoleInsert(authRole *AuthRole) error {
	data := make(map[string]interface{})
	if authRole.Name != "" {
		data["name"] = authRole.Name
	}
	if authRole.Pid > 0 {
		data["pid"] = authRole.Pid
	}
	if authRole.Status >= 0 {
		data["status"] = authRole.Status
	}
	if authRole.Remark != "" {
		data["remark"] = authRole.Remark
	}
	if authRole.Listorder > 0 {
		data["listorder"] = authRole.Listorder
	}
	data["create_time"] = time.Now()
	data["update_time"] = time.Now()
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthRole{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthRoleUpdate(authRole *AuthRole) error {
	data := make(map[string]interface{})
	if authRole.Name != "" {
		data["name"] = authRole.Name
	}
	if authRole.Pid > 0 {
		data["pid"] = authRole.Pid
	}
	if authRole.Status >= 0 {
		data["status"] = authRole.Status
	}
	if authRole.Remark != "" {
		data["remark"] = authRole.Remark
	}
	if authRole.Listorder > 0 {
		data["listorder"] = authRole.Listorder
	}
	data["update_time"] = time.Now()
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthRole{}).Where("id = ?", authRole.Id).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthRoleDeleteById(id int64) error {
	err := DB.Where("id = ?", id).Delete(&AuthRole{}).Error
	if err != nil {
		return err
	}
	return nil
}
