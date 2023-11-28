package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
)

func main() {
	conf.Init()
	h := routes.Router()
	h.Spin()
}
