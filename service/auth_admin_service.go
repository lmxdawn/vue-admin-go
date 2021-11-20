package service

import (
	"fmt"
	"github.com/lmxdawn/vue-admin-go/model"
	"github.com/lmxdawn/vue-admin-go/req"
	"github.com/lmxdawn/vue-admin-go/res"
)

type AuthAdminService struct {
}

func (s *AuthAdminService) ListAdminPage(q req.AuthAdminQueryReq) (*res.PageSimpleRes, error) {
	pageRes := &res.PageSimpleRes{}
	fmt.Println("1111", pageRes)
	total, list, err := model.AuthAdminListPage(q)
	if err != nil {
		return nil, err
	}
	fmt.Println("22222")
	pageRes.Total = total
	pageRes.List = list
	return pageRes, err
}
