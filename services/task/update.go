package task

import (
	"net/http"

	"github.com/alifudin-a/go-task-tracker/helpers"
	"github.com/alifudin-a/go-task-tracker/models"
	repo "github.com/alifudin-a/go-task-tracker/repo/task"
	"github.com/labstack/echo/v4"
)

func UpdateTask(c echo.Context) (err error) {
	var resp helpers.Response
	var reqTask = new(models.Task)

	if err = c.Bind(&reqTask); err != nil {
		return
	}

	taskRepo := repo.NewTaskRepository()

	arg := repo.UpdateTaskArgument{
		ID:       reqTask.ID,
		Text:     reqTask.Text,
		Day:      reqTask.Day,
		Reminder: reqTask.Reminder,
	}

	var task *models.Task

	task, err = taskRepo.Update(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully Update Task!"
	resp.Body = map[string]interface{}{
		"task": task,
	}

	return c.JSON(http.StatusOK, task)
}
