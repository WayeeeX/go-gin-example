package request

type PageQuery struct {
	PageSize int    `form:"page_size" binding:"required,number,min=1"'`
	PageNum  int    `form:"page_num" binding:"required,number,min=1"`
	Keyword  string `form:"keyword"`
}
type UpdateStatus struct {
	IdsJson
	Status *int `json:"status" binding:"required,number,oneof=-1 0 1"`
}

type IdsJson struct {
	Ids []int64 `json:"ids" binding:"required,min=1"`
}

type IdQuery struct {
	Id int64 `form:"id" binding:"required,number,min=1"`
}
