package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type TodosResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func GetAllTodoController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos", func(ctx echo.Context) error {
		query := "SELECT id, title, description, done FROM todos"
		rows, err := db.Query(query)

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
