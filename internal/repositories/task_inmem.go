package repositories

import (
	"sync"

	"github.com/andrewesteves/tasks/internal/entities"
)

type TaskInMemory struct {
	sync.RWMutex
	db map[string]entities.Task
}

func NewTaskInMemory() Task {
	return &TaskInMemory{
		db: make(map[string]entities.Task),
	}
}

func (t *TaskInMemory) Get() []entities.Task {
	t.RLock()
	defer t.RUnlock()
	var tasks []entities.Task

	for key := range t.db {
		tasks = append(tasks, t.db[key])
	}

	return tasks
}

func (t *TaskInMemory) Put(task entities.Task) {
	t.Lock()
	defer t.Unlock()
	t.db[task.ID] = task
}

func (t *TaskInMemory) Delete(id string) {
	t.Lock()
	defer t.Unlock()
	delete(t.db, id)
}
