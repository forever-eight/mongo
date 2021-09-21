package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/ds"
)

func (e *Endpoint) find(c echo.Context) error {
	var input ds.Project
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with unmarshalling")
	}

	found, err := e.s.Find(input.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with delete")
	}
	return c.JSON(http.StatusOK, found)
}