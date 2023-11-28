package handler

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

var identityKey = "id"

func IdLogin(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	identity, _ := c.Get(identityKey)
	var user model.UserDao
	count := 0
	Dao.DB.Model(&user).Where("user_name=?", identity.(*service.UserService).UserName).First(&user).Count(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: e.ErrorNotExistUser,
			Msg:    id,
		})
	} else {
		if id == strconv.Itoa(int(user.ID)) {
			c.JSON(http.StatusBadRequest, model.Response{
				Status: e.SUCCESS,
				Data:   model.BuildUser(user),
				Msg:    "登录成功",
			})
		} else {
			c.JSON(http.StatusBadRequest, model.Response{
				Status: e.ErrorId,
				Msg:    "用户id错误",
			})
		}
	}

	//user, _ := c.Get(identityKey)
	//c.JSON(200, utils.H{
	//	"message": fmt.Sprintf("username:%v", user.(*service.UserService).UserName),
	//})
}
