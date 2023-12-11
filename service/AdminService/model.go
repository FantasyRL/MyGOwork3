package AdminService

type ListUsersService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type AddAdminService struct {
}
type BlockUserService struct {
}
