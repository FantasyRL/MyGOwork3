package userHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary UserLogin
// @Description Login
// @Accept json/form
// @Produce json
// @Param user_name query string true "用户名" minlength(4) maxlength(15)
// @Param password query string true "密码" minlength(6) maxlength(32)
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Failure 500 {object} model.ErrorResponse "内部错误"
// @Router /user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) (*UserService.UserService, error, model.Response) {
	var userLogin UserService.UserService
	if err := c.BindAndValidate(&userLogin); err == nil { //相当于gin shouldBind 将请求携带的参数和后端的结构体绑定起来
		res := userLogin.Login()
		return &userLogin, nil, res
	} else {
		return &userLogin, err, model.Response{
			Status: e.InvalidParams,
			Msg:    "参数出错",
		}
	}
}
