package main

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"yourapp/api/hello"
	"yourapp/cache"
	"yourapp/common"
	"yourapp/database"
	"yourapp/middleware"
	"yourapp/route"
)

func main() {
	// Load .env file at the very first
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize software compenents
	// init logging system
	common.InitLogging()
	logger := common.GetLogger()

	common.GetLoggerWithAction("StartService").Info("version v1.0")

	// init database connnection
	err = database.Init()
	if err != nil {
		common.GetLoggerWithAction("InitMySQL").Fatal(err.Error())
	} else {
		common.GetLoggerWithAction("InitMySQL").Info("MySQL connected")
	}

	err = cache.Init()
	if err != nil {
		common.GetLoggerWithAction("InitRedis").Fatal(err.Error())
	} else {
		common.GetLoggerWithAction("InitRedis").Info("Redis connected")
	}

	// init cron jobs
	// c := cron.New(cron.WithSeconds())
	// c.AddFunc("*/5 * * * * *", func() { fmt.Println("Every 5s") })
	// c.Start()

	// init web server
	common.BodyValidator = validator.New()

	e := echo.New()

	controllers := []common.Controller{
		hello.Controller{},
	}

	e.Use(middleware.LoggingHeader)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			//return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})
	s := NewStats()
	e.Use(s.Process)

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	route.DefineAPIRoutes(e, "/api/v0", controllers)

	port := ":" + os.Getenv("RESTAPI_PORT")
	logger.WithField("action", "EchoServer").Fatal(e.Start(port))

}

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}
