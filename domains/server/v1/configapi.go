package serverv1

import (
	// stdlib
	"net/http"

	// other
	"github.com/labstack/echo"
)

// This function responsible for getting runtime configuration.
func configurationGET(ec echo.Context) error {
	return ec.JSON(http.StatusOK, map[string]string{"result": "success"})
}
