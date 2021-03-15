package main

import (
	"net/http"

	database "github.com/alifudin-a/go-task-tracker/database/psql"
	"github.com/alifudin-a/go-task-tracker/services/task"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.DBConn()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, URI=${uri}, STATUS=${status}, " +
			"HOST=${host}, ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("api")
	api.GET("/tasks", task.ListTasks)
	api.GET("/task/:id", task.ReadTask)
	api.DELETE("/task/:id", task.DeleteTask)
	api.POST("/task", task.CreateTask)
	api.PUT("/task", task.UpdateTask)

	e.Logger.Fatal(e.Start(":5000"))
}
