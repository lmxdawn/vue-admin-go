package model

import (
	"fmt"
	"github.com/lmxdawn/vue-admin-go/req"
	"time"
)

type AuthAdmin struct {
	Id            int        //
	Username      string     // 用户名
	Password      string     // 登录密码；sp_password加密
	Tel           string     // 用户手机号
	Email         string     // 登录邮箱
	Avatar        string     // 用户头像
	Sex           int16      // 性别；0：保密，1：男；2：女
	LastLoginIp   string     // 最后登录ip
	LastLoginTime *time.Time // 最后登录时间
	CreateTime    *time.Time // 注册时间
	Status        int        // 用户状态 0：禁用； 1：正常 ；2：未验证
}

func AuthAdminListPage(q req.AuthAdminQueryReq) (int64, []AuthAdmin, error) {
	var authAdmins []AuthAdmin
	offset := (q.Page - 1) * q.Limit
	fmt.Println("33333", DB)
	db := DB.Select("id,username,avatar,tel,email,status,last_login_ip,last_login_time,create_time")
	if q.Ids != nil && len(q.Ids) != 0 {
		db = db.Where("id IN ? ", q.Ids)
	}
	if q.Status != -1 {
		db = db.Where("status = ? ", q.Status)
	}
	if q.Username != "" {
		db = db.Where("username LIKE ? ", fmt.Sprint("%", q.Username, "%"))
	}
	err := db.Offset(offset).Limit(q.Limit).Order("id DESC").Find(&authAdmins).Error
	if err != nil {
		return 0, authAdmins, err
	}
	var total int64
	db.Count(&total)
	return 0, authAdmins, nil
}

func AuthAdminFindByUserName(username string) (*AuthAdmin, error) {
	var authAdmin *AuthAdmin
	db := DB.Select("id,password")
	err := db.Where("username = ?", username).First(&authAdmin).Error
	if err != nil {
		return nil, err
	}
	return authAdmin, nil
}

func AuthAdminFindById(id int64) (*AuthAdmin, error) {
	var authAdmin *AuthAdmin
	db := DB.Select("id,username,avatar")
	err := db.Where("id = ?", id).First(&authAdmin).Error
	if err != nil {
		return nil, err
	}
	return authAdmin, nil
}

func AuthAdminFindPwdById(id int64) (*AuthAdmin, error) {
	var authAdmin *AuthAdmin
	db := DB.Select("password")
	err := db.Where("id = ?", id).First(&authAdmin).Error
	if err != nil {
		return nil, err
	}
	return authAdmin, nil
}

func AuthAdminInsert(authAdmin *AuthAdmin) error {
	data := make(map[string]interface{})
	if authAdmin.Password != "" {
		data["password"] = authAdmin.Password
	}
	if authAdmin.Tel != "" {
		data["tel"] = authAdmin.Tel
	}
	if authAdmin.Email != "" {
		data["email"] = authAdmin.Email
	}
	if authAdmin.Avatar != "" {
		data["avatar"] = authAdmin.Avatar
	}
	if authAdmin.Sex > 0 {
		data["sex"] = authAdmin.Sex
	}
	data["username"] = authAdmin.Username
	data["create_time"] = time.Now()
	data["status"] = 1
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthAdmin{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthAdminUpdate(authAdmin *AuthAdmin) error {
	data := make(map[string]interface{})
	if authAdmin.Username != "" {
		data["username"] = authAdmin.Username
	}
	if authAdmin.Password != "" {
		data["password"] = authAdmin.Password
	}
	if authAdmin.Tel != "" {
		data["tel"] = authAdmin.Tel
	}
	if authAdmin.Email != "" {
		data["email"] = authAdmin.Email
	}
	if authAdmin.Avatar != "" {
		data["avatar"] = authAdmin.Avatar
	}
	if authAdmin.Sex > 0 {
		data["sex"] = authAdmin.Sex
	}
	if authAdmin.LastLoginIp != "" {
		data["last_login_ip"] = authAdmin.LastLoginIp
	}
	if authAdmin.LastLoginTime != nil {
		data["last_login_time"] = authAdmin.LastLoginTime
	}
	if authAdmin.Status >= 0 {
		data["status"] = authAdmin.Status
	}
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthAdmin{}).Where("id = ?", authAdmin.Id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthAdminDeleteById(id int64) error {
	err := DB.Where("id = ?", id).Delete(&AuthAdmin{}).Error
	if err != nil {
		return err
	}
	return nil
}
