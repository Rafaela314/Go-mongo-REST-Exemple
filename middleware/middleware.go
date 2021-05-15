package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Logger returns a middleware that logs HTTP requests.
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			log.WithFields(log.Fields{
				"id":          id,
				"ip":          c.RealIP(),
				"finished":    stop.Format(time.RFC3339),
				"host":        req.Host,
				"method":      req.Method,
				"requestURI":  req.RequestURI,
				"status":      res.Status,
				"requestSize": strconv.FormatInt(res.Size, 10),
				"elapsed":     stop.Sub(start).String(),
				"referer":     req.Referer(),
				"userAgent":   req.UserAgent(),
			}).Info("completed handling request")
			return nil
		}
	}
}

// CustomContextMiddelware sets pairs 'key:value' in the default echo.Context
func CustomContextMiddelware(key string, val interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(key, val)
			return next(c)
		}

	}
}
