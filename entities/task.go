package entities

type Task struct {
	Id       int64
	Task     string
	Asignee  string
	Deadline string
	Status   string
}
