package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/ds"
)

func (e *Endpoint) delete(c echo.Context) error {
	var input ds.Project
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with unmarshalling")
	}

	err = e.s.Delete(input.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with delete")
	}
	return c.String(http.StatusOK, "delete ok")
}
