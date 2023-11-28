package routes

import (
	"ToDoList/handler"
	jwt "ToDoList/pkg/authorization"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func Router() *server.Hertz {
	//session扩展
	h := server.New(server.WithHostPorts(":8000"))
	h.Use(accesslog.New())
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("PANDORA PARADOXXX"))
	store := cookie.NewStore([]byte("PANDORA PARADOXXX"))
	h.Use(sessions.New("sessionId", store))
	mw := jwt.JWT()
	v1 := h.Group("todolist/")
	{
		v1.POST("user/register", handler.UserRegister)
		v1.POST("user/login", mw.LoginHandler) //Hertz JWT:Authenticator
	}
	//尝试实现带着id来访问
	auth := v1.Group("/auth")
	auth.Use(mw.MiddlewareFunc())
	{
		auth.GET("/test", jwt.TestHandler)
		auth.GET("/:id", handler.IdLogin)
	}

	return h
}
