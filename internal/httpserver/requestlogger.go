package httpserver

import (
	// other

	"time"

	"github.com/labstack/echo"
)

func requestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			startTime := time.Now()

			err := next(ec)

			log.Info().
				Str("From", ec.RealIP()).
				Str("To", ec.Request().Host).
				Str("Method", ec.Request().Method).
				Str("Path", ec.Request().URL.Path).
				Int64("Length", ec.Request().ContentLength).
				Str("UA", ec.Request().UserAgent()).
				TimeDiff("TimeMS", time.Now(), startTime).
				Msg("HTTP request")

			return err
		}
	}
}
