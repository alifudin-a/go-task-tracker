package task

import (
	"log"
	"net/http"

	"github.com/alifudin-a/go-task-tracker/helpers"
	"github.com/alifudin-a/go-task-tracker/models"
	repo "github.com/alifudin-a/go-task-tracker/repo/task"
	"github.com/labstack/echo/v4"
)

func ListTasks(c echo.Context) (err error) {

	var tasks []models.Task
	var resp helpers.Response

	taskRepo := repo.NewTaskRepository()

	tasks, err = taskRepo.FindAll()
	if err != nil {
		log.Println("An error occured", err)
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully List All Task!"
	resp.Body = map[string]interface{}{
		"tasks": tasks,
	}

	return c.JSON(http.StatusOK, resp)
}
