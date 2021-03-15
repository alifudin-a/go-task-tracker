package task

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-task-tracker/helpers"
	repo "github.com/alifudin-a/go-task-tracker/repo/task"
	"github.com/labstack/echo/v4"
)

func DeleteTask(c echo.Context) (err error) {
	var resp helpers.Response
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	taskRepo := repo.NewTaskRepository()

	arg := repo.DeleteTaskArgument{
		ID: int64(id),
	}

	err = taskRepo.Delete(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully Delete Task!"

	return c.JSON(http.StatusOK, resp)
}
