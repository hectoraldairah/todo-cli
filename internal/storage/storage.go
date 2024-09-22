package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hectoraldairah/todo-cli/internal/task"
)

var fileName = "tasks.json"

func LoadTask() ([]task.Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("File does not exists, new file created\n")
		return []task.Task{}, nil
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var tasks []task.Task

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTask(tasks []task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
