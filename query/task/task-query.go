package query

var ListTasks = `SELECT * FROM task;`

var GetTask = `SELECT * FROM task WHERE id = $1;`

var DeleteTask = `DELETE FROM task WHERE id =$1;`

var UpdateTask = `UPDATE task SET text = $1, day = $2, reminder = $3 WHERE id = $4 RETURNING *;`

var CreateTask = `INSERT INTO task (text, day, reminder) VALUES ($1, $2, $3) RETURNING*;`
