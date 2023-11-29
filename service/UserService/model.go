package UserService

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=4,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}
