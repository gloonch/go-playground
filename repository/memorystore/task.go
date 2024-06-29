package memorystore

import (
	"fmt"
	"todocli/entity"
)

type Task struct {
	tasks []entity.Task
}

func NewTaskStore() *Task {
	return &Task{
		tasks: make([]entity.Task, 0),
	}
}

// we have to put pointer because we change the Task state in this method
func (t *Task) CreateNewTask(task entity.Task) (entity.Task, error) {
	task.ID = len(t.tasks) + 1

	t.tasks = append(t.tasks, task)

	return task, nil
}

func (t Task) ListUserTasks(userID int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, task := range t.tasks {
		if task.UserID == userID {
			fmt.Println(task)
		}
	}
	return userTasks, nil
}
