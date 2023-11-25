package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteTodoController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		query := "DELETE FROM todos WHERE id = ?"
		_, err := db.Exec(query, id)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Data success deleted")

	})
}
