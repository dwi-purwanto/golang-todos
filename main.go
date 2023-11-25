package main

import (
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	controller.GetAllTodoController(e, db)
	controller.CreateTodoController(e, db)
	controller.UpdateTodoController(e, db)
	controller.UpdateCheckController(e, db)
	controller.ShowTodoController(e, db)
	controller.DeleteTodoController(e, db)

	e.Start(":8000")

}
