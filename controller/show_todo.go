package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func ShowTodoController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		query := "SELECT id, title, description, done FROM todos WHERE id = ?"
		rows, err := db.Query(query, id) // TODO: Add error handling
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		var res []TodosResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int
			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
			var todo TodosResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}
			res = append(res, todo)
		}
		return ctx.JSON(http.StatusOK, res)
	})
}
