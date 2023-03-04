package repositories

import "github.com/andrewesteves/tasks/internal/entities"

type Task interface {
	TaskReader
	TaskWriter
}

type TaskReader interface {
	Get() []entities.Task
}

type TaskWriter interface {
	Put(task entities.Task)
	Delete(id string)
}
