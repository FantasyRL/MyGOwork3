package adminHandler

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

var identityKey = "id"

func adminLogin(ctx context.Context, c *app.RequestContext) bool {
	var user model.UserDao
	identity, _ := c.Get(identityKey)
	count := 0
	Dao.DB.Model(&user).Where("user_name=?", identity.(*UserService.UserService).UserName).First(&user).Count(&count)
	if count == 0 {
		return false
	} else {
		if user.Status == 114514 {
			return true
		} else {
			return false
		}
	}
}
