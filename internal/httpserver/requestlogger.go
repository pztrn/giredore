package httpserver

import (
	"time"

	"github.com/labstack/echo"
)

func requestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			startTime := time.Now()

			err := next(ectx)

			log.Info().
				Str("From", ectx.RealIP()).
				Str("To", ectx.Request().Host).
				Str("Method", ectx.Request().Method).
				Str("Path", ectx.Request().URL.Path).
				Int64("Length", ectx.Request().ContentLength).
				Str("UA", ectx.Request().UserAgent()).
				TimeDiff("TimeMS", time.Now(), startTime).
				Msg("HTTP request")

			return err
		}
	}
}
