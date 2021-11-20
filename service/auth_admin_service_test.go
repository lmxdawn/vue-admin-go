package service

import (
	"fmt"
	"github.com/lmxdawn/vue-admin-go/req"
	"testing"
)

func TestListAdminPage(t *testing.T) {
	q := req.AuthAdminQueryReq{}
	q.Page = 1
	q.Limit = 20
	fmt.Println(q)
	a := &AuthAdminService{}
	res, err := a.ListAdminPage(q)
	if err != nil {

	}
	fmt.Println(res)
}
