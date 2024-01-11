package handlers

import (
	"database/sql"
	dbc "echo-test/database/queries"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Query(c echo.Context) error {
	db := c.Get("db").(*sql.DB)

	results, err := dbc.GetMarks(db)
	if err != nil {
		return c.String(http.StatusFailedDependency, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}
