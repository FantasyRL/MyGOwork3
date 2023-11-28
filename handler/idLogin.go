package handler

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

var identityKey = "id"

func IdLogin(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	identity, _ := c.Get(identityKey)
	var user model.UserDao
	count := 0
	Dao.DB.Model(&user).Where("user_name=?", identity.(*service.UserService).UserName).First(&user).Count(&count)
	if count == 0 {
		c.JSON(http.StatusOK, model.Response{
			Status: e.ErrorNotExistUser,
			Msg:    id,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Status: e.SUCCESS,
			Msg:    id,
		})
	}

	//user, _ := c.Get(identityKey)
	//c.JSON(200, utils.H{
	//	"message": fmt.Sprintf("username:%v", user.(*service.UserService).UserName),
	//})
}
