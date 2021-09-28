package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/ds"
)

func (e *Endpoint) find(c echo.Context) error {
	var input ds.JustID
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	fmt.Println(input)
	if err != nil {
		fmt.Println(input)
		return c.String(http.StatusBadRequest, "problem with unmarshalling")
	}
	found, err := e.s.Find(input.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with find"+err.Error())
	}
	return c.JSON(http.StatusOK, found)
}
