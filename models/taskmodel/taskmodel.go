package taskmodel

import (
	"database/sql"
	"tasking-app-golang/config"
	"tasking-app-golang/entities"
)

type TaskModel struct {
	db *sql.DB
}

func New() *TaskModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &TaskModel{db: db}
}

func (m *TaskModel) FindAll(task *[]entities.Task) error {
	rows, err := m.db.Query("select * from task")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Task
		rows.Scan(
			&data.Id,
			&data.Task,
			&data.Asignee,
			&data.Deadline,
			&data.Status)

		*task = append(*task, data)
	}

	return nil
}

func (m *TaskModel) Create(task *entities.Task) error {
	result, err := m.db.Exec("insert into task (task, asignee, deadline) values(?,?,?)",
		task.Task, task.Asignee, task.Deadline)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	task.Id = lastInsertId
	return nil
}

func (m *TaskModel) Find(id int64, task *entities.Task) error {
	return m.db.QueryRow("select * from task where id = ?", id).Scan(
		&task.Id,
		&task.Task,
		&task.Asignee,
		&task.Deadline,
		&task.Status)
}

func (m *TaskModel) Update(task entities.Task) error {

	_, err := m.db.Exec("update task set task = ?, asignee = ?, deadline = ? where id = ?",
		task.Task, task.Asignee, task.Deadline, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) UpdateTask(id int64) error {

	_, err := m.db.Exec("update task set status = 'Done' where id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from task where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
