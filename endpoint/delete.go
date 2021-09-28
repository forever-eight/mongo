package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/ds"
)

func (e *Endpoint) delete(c echo.Context) error {
	var input ds.JustID
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	fmt.Println(input)
	if err != nil {
		fmt.Println(input)
		return c.String(http.StatusBadRequest, "problem with unmarshalling")
	}

	err = e.s.Delete(input.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with delete "+err.Error())
	}
	return c.String(http.StatusOK, "delete ok")
}
