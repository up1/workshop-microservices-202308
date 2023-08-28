package demo

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func SayHi(connection *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		data, _ := GetData(connection)
		res := HelloResponse{Message: data}
		return c.JSON(http.StatusOK, res)
	}
}

// Get data from database
func GetData(connection *pgxpool.Pool) (string, error) {
	rows, err := connection.Query(context.Background(), "select * from message")
	if err != nil {
		return "", err
	}

	// iterate through the rows
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return "", err
		}

		// convert DB types to Go types
		data := values[1].(string)
		return data, nil
	}
	return "", nil
}
