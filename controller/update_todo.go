package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		query := "UPDATE todos SET title = ?, description = ? WHERE id = ?"
		_, err := db.Exec(query,
			request.Title,
			request.Description,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Data success updated")
	})
}
