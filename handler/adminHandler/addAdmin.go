package adminHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/AdminService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary AddAdmin
// @Description AddAdmin
// @Accept json/form
// @Produce json
// @Param id path string true "用户id"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Router /admin/add/{id} [GET]
func AddAdmin(ctx context.Context, c *app.RequestContext) {
	if !adminLogin(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorAuth,
			Msg:    "你没有访问权限",
		})
		return
	}
	id := c.Param("id")
	var addAdmin AdminService.AddAdminService
	res := addAdmin.AddAdmin(id)
	c.JSON(200, res)
}
