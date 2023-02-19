package request

type PageQuery struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}
