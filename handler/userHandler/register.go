package userHandler

import (
	"ToDoList/pkg/e"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary UserRegister
// @Description Register
// @Accept json/form
// @Produce json
// @Param user_name query string true "用户名" minlength(4) maxlength(15)
// @Param password query string true "密码" minlength(6) maxlength(32)
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Router /user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var userRegister UserService.UserService
	if err := c.BindAndValidate(&userRegister); err == nil { //相当于gin shouldBind 将请求携带的参数和后端的结构体绑定起来
		res := userRegister.Register()
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, err)
	}
}
