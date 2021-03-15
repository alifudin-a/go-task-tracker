package task

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-task-tracker/helpers"
	"github.com/alifudin-a/go-task-tracker/models"
	repo "github.com/alifudin-a/go-task-tracker/repo/task"
	"github.com/labstack/echo/v4"
)

func ReadTask(c echo.Context) (err error) {
	var resp helpers.Response
	task := &models.Task{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	taskRepo := repo.NewTaskRepository()

	arg := repo.FindByIdArgument{
		ID: int64(id),
	}

	task, err = taskRepo.FindById(arg)
	if err != nil {
		log.Println(err)
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully Get Task!"
	resp.Body = map[string]interface{}{
		"task": task,
	}

	return c.JSON(http.StatusOK, resp)
}
