package hello

import (
	"net/http"
	"yourapp/common"

	"github.com/labstack/echo/v4"
)

type HelloBody struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// Controller holds the routes for vehicle controller
type Controller struct {
}

// Routes returns all the routes of the vehicle controller
func (c Controller) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.GET,
			Path:    "/hello",
			Handler: c.HelloHandler,
		},
	}
}

func (c Controller) HelloHandler(ctx echo.Context) error {
	logger := common.GetLoggerWithAction("HelloHandler")
	resp := common.NewReponse()
	body := common.APIResponseBody{}

	logger.Warn("Hello Handler is called")

	body["data"] = HelloBody{ID: "1", Name: "yourapp"}

	resp.SetBody(body)
	return ctx.JSON(http.StatusOK, resp.Build())
}
