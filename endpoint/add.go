package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/ds"
)

func (e *Endpoint) add(c echo.Context) error {
	var input ds.Project
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with unmarshalling")
	}

	err = e.s.Add(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with add "+err.Error())
	}
	return c.String(http.StatusOK, "add ok")
}
