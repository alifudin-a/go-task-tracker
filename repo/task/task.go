package repo

import (
	"log"

	"github.com/alifudin-a/go-task-tracker/database/psql"
	"github.com/alifudin-a/go-task-tracker/models"
	query "github.com/alifudin-a/go-task-tracker/query/task"
)

type TaskRepository interface {
	FindAll() ([]models.Task, error)
	FindById(arg FindByIdArgument) (*models.Task, error)
	Delete(arg DeleteTaskArgument) (err error)
	Create(arg CreateTaskArgument) (*models.Task, error)
	Update(arg UpdateTaskArgument) (*models.Task, error)
}

type repo struct{}

func NewTaskRepository() TaskRepository {
	return &repo{}
}

type FindByIdArgument struct {
	ID int64 `json:"id"`
}

func (*repo) FindById(arg FindByIdArgument) (*models.Task, error) {
	var task models.Task
	var db = psql.DBConn()

	err := db.Get(&task, query.GetTask, arg.ID)
	if err != nil {
		log.Println("An error occured when get task: ", err)
	}

	return &task, nil
}

func (*repo) FindAll() ([]models.Task, error) {
	var task []models.Task
	db := psql.DBConn()

	err := db.Select(&task, query.ListTasks)
	if err != nil {
		log.Println("An error occured while select", err)
	}

	return task, nil
}

type DeleteTaskArgument struct {
	ID int64 `json:"id"`
}

func (*repo) Delete(arg DeleteTaskArgument) (err error) {
	db := psql.DBConn()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteTask, arg.ID)
	if err != nil {
		tx.Rollback()
		log.Println("An error occured while deleting task: ", err)
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return nil
}

type CreateTaskArgument struct {
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func (*repo) Create(arg CreateTaskArgument) (*models.Task, error) {
	db := psql.DBConn()
	var task models.Task

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateTask, arg.Text, arg.Day, arg.Reminder).StructScan(&task)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &task, nil
}

type UpdateTaskArgument struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func (*repo) Update(arg UpdateTaskArgument) (*models.Task, error) {
	db := psql.DBConn()
	var task models.Task

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateTask, arg.Text, arg.Day, arg.Reminder, arg.ID).StructScan(&task)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &task, nil
}
