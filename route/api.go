package route

import (
	"yourapp/common"

	"github.com/labstack/echo/v4"
)

// DefineAPIRoutes adds routes to server, add controller to the controller in the function in order
// to add more routes of controller
func DefineAPIRoutes(e *echo.Echo, prefix string, controllers []common.Controller) {
	var routes []common.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	api := e.Group(prefix)
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
