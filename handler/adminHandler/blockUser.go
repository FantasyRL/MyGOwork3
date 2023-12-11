package adminHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/AdminService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary BlockUser
// @Description BlockUser
// @Accept json/form
// @Produce json
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} e.InvalidParams "请求错误"
// @Failure 500 {object} e.ERROR "内部错误"
// @Router /admin/block/:id [GET]
func BlockUser(ctx context.Context, c *app.RequestContext) {
	if !adminLogin(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorAuth,
			Msg:    "你没有访问权限",
		})
		return
	}
	id := c.Param("id")
	var blockUser AdminService.BlockUserService
	res := blockUser.Block(id)
	c.JSON(200, res)
}
