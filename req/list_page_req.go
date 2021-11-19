package req

type ListPage struct {
	Page  int `json:"page" binding:"required,gte=1"`      // 页数
	Limit int `json:"limit" binding:"required,gte=1,lte"` // 每页返回多少
}
