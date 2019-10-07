package serverv1

import (
	// stdlib
	"net/http"

	// other
	"github.com/labstack/echo"
)

func throwGoImports(ec echo.Context) error {
	return ec.String(http.StatusOK, "All OK here")
}
