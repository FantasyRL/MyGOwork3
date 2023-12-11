package userHandler

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

var identityKey = "id"

// @Summary idLogin test
// @Description idLogin
// @Accept json/form
// @Produce json
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} e.InvalidParams "请求错误"
// @Failure 500 {object} e.ERROR "内部错误"
// @Router /auth/:id [GET]
func IdLogin(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	identity, _ := c.Get(identityKey)
	var user model.UserDao
	count := 0
	Dao.DB.Model(&user).Where("user_name=?", identity.(*UserService.UserService).UserName).First(&user).Count(&count)
	if count == 0 {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorNotExistUser,
			Msg:    id,
		})
	} else {
		if id == strconv.Itoa(int(user.ID)) {
			c.JSON(e.InvalidParams, model.Response{
				Status: e.SUCCESS,
				Data:   model.BuildUser(user),
				Msg:    "登录成功",
			})
		} else {
			c.JSON(e.InvalidParams, model.ErrorResponse{
				Status: e.ErrorId,
				Msg:    "用户id错误",
			})
		}
	}
}

func IdTask(ctx context.Context, c *app.RequestContext) bool {
	id := c.Param("id")
	identity, _ := c.Get(identityKey)
	var user model.UserDao
	count := 0
	Dao.DB.Model(&user).Where("user_name=?", identity.(*UserService.UserService).UserName).First(&user).Count(&count)
	if count == 0 || user.Status == 1 {
		return false
	} else {
		if id == strconv.Itoa(int(user.ID)) {
			return true
		} else {
			return false
		}
	}
}
