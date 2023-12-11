package AdminService

type ListUsers struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type AddAdmin struct {
}
