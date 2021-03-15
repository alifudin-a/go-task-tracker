package task

import (
	"net/http"

	"github.com/alifudin-a/go-task-tracker/helpers"
	"github.com/alifudin-a/go-task-tracker/models"
	repo "github.com/alifudin-a/go-task-tracker/repo/task"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) (err error) {
	var resp helpers.Response
	var reqTask = new(models.Task)
	var task *models.Task

	if err = c.Bind(&reqTask); err != nil {
		return
	}

	taskRepo := repo.NewTaskRepository()

	arg := repo.CreateTaskArgument{
		Text:     reqTask.Text,
		Day:      reqTask.Day,
		Reminder: reqTask.Reminder,
	}

	task, err = taskRepo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Successfully Create Task!"
	resp.Body = map[string]interface{}{
		"task": task,
	}

	return c.JSON(http.StatusCreated, resp)
}
