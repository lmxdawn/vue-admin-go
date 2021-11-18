package model

import "time"

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
