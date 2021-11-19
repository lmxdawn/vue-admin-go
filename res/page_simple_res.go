package res

type PageSimpleRes struct {
	Total int64         `json:"total"` // 分页总数
	List  []interface{} `json:"list"`  // 分页列表
}
