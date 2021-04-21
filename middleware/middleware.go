package middleware

import (
	"yourapp/common"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// LoggingHeader is called by request beign routed
// echo Middleware 是一个 echo HandlerFunc, 会在request被处理前调用
func LoggingHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := common.GetLogger().WithFields(logrus.Fields{
			"action": "Echo",
			"method": c.Request().Method,
			"url":    c.Request().RequestURI,
		})
		logger.Info("remote address:", c.Request().RemoteAddr)
		return next(c)
	}
}
