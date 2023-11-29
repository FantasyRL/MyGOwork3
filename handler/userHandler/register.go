package userHandler

import (
	"ToDoList/pkg/e"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var userRegister UserService.UserService
	if err := c.BindAndValidate(&userRegister); err == nil { //相当于gin shouldBind 将请求携带的参数和后端的结构体绑定起来
		res := userRegister.Register()
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, err)
	}
}
