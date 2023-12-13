package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
	_ "github.com/hertz-contrib/swagger/example/basic/docs"
)

// @title           TodoList
// @version         1.0
// @description     CRUD your todoList

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /todolist

func main() {
	conf.Init()
	h := routes.Router()
	h.Spin()
}
