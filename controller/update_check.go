package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type CheckRequest struct {
	Done bool `json:"done"`
}

func UpdateCheckController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var done int
		if request.Done {
			done = 1
		}
		query := "UPDATE todos SET done = ? WHERE id = ?"
		_, err := db.Exec(query,
			done,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Data success updated")
	})
}
