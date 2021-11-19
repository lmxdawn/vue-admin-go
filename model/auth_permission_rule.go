package model

import "time"

type AuthPermissionRule struct {
	Id         int        // 规则编号
	Pid        int        // 父级id
	Name       string     // 规则唯一标识
	Title      string     // 规则中文名称
	Status     int        // 状态：为1正常，为0禁用
	Condition  string     // 规则表达式，为空表示存在就验证，不为空表示按照条件验证
	Listorder  int        // 排序，优先级，越小优先级越高
	CreateTime *time.Time // 创建时间
	UpdateTime *time.Time // 更新时间
}

func AuthPermissionRuleListByIdIn(ids []int64) ([]AuthPermissionRule, error) {
	var authPermissionRules []AuthPermissionRule
	db := DB.Select("name")
	err := db.Where("id IN ?", ids).Find(&authPermissionRules).Error
	if err != nil {
		return nil, err
	}
	return authPermissionRules, nil
}

func AuthPermissionRuleListByPid(pid int64) ([]AuthPermissionRule, error) {
	var authPermissionRules []AuthPermissionRule
	db := DB.Select("`id`,`pid`,`name`,`title`,`status`,`condition`,`listorder`")
	err := db.Where("pid = ?", pid).Order("listorder DESC").Find(&authPermissionRules).Error
	if err != nil {
		return nil, err
	}
	return authPermissionRules, nil
}

func AuthPermissionRuleListAll() (*[]AuthPermissionRule, error) {
	var authPermissionRules *[]AuthPermissionRule
	db := DB.Select("`id`,`pid`,`name`,`title`,`status`,`condition`,`listorder`")
	err := db.Find(authPermissionRules).Error
	if err != nil {
		return nil, err
	}
	return authPermissionRules, nil
}

func AuthPermissionRuleFindByName(name string) ([]AuthPermissionRule, error) {
	var authPermissionRules []AuthPermissionRule
	db := DB.Select("`id`")
	err := db.Where("name = ?", name).Find(&authPermissionRules).Error
	if err != nil {
		return nil, err
	}
	return authPermissionRules, nil
}

func AuthPermissionRuleInsert(authPermissionRule *AuthPermissionRule) error {
	data := make(map[string]interface{})
	if authPermissionRule.Pid > 0 {
		data["pid"] = authPermissionRule.Pid
	}
	if authPermissionRule.Name != "" {
		data["name"] = authPermissionRule.Name
	}
	if authPermissionRule.Title != "" {
		data["title"] = authPermissionRule.Title
	}
	if authPermissionRule.Status >= 0 {
		data["status"] = authPermissionRule.Status
	}
	if authPermissionRule.Condition != "" {
		data["condition"] = authPermissionRule.Condition
	}
	if authPermissionRule.Listorder > 0 {
		data["listorder"] = authPermissionRule.Listorder
	}
	data["create_time"] = time.Now()
	data["update_time"] = time.Now()
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthPermissionRule{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthPermissionRuleUpdate(authPermissionRule *AuthPermissionRule) error {
	data := make(map[string]interface{})
	if authPermissionRule.Pid > 0 {
		data["pid"] = authPermissionRule.Pid
	}
	if authPermissionRule.Name != "" {
		data["name"] = authPermissionRule.Name
	}
	if authPermissionRule.Title != "" {
		data["title"] = authPermissionRule.Title
	}
	if authPermissionRule.Status >= 0 {
		data["status"] = authPermissionRule.Status
	}
	if authPermissionRule.Condition != "" {
		data["condition"] = authPermissionRule.Condition
	}
	if authPermissionRule.Listorder > 0 {
		data["listorder"] = authPermissionRule.Listorder
	}
	data["update_time"] = time.Now()
	if data == nil {
		return nil
	}
	err := DB.Model(&AuthPermissionRule{}).Where("id = ", authPermissionRule.Id).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthPermissionRuleDeleteById(id int64) error {
	err := DB.Where("id = ?", id).Delete(&AuthPermissionRule{}).Error
	if err != nil {
		return err
	}
	return nil
}
