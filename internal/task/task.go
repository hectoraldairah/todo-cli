package task

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/hectoraldairah/todo-cli/internal/storage"
)

type Status string

type Task struct {
	ID     int
	Name   string
	Status Status
}

const (
	ToDo       Status = "To do"
	InProgress Status = "In progress"
	Done       Status = "Done"
)

func AddTask(description string) error {
	_, err := storage.GetDB().Exec("INSERT INTO tasks (description, status) values (?, ?)", description, Done)
	return err
}

func GetTasks() ([]Task, error) {
	rows, err := storage.GetDB().Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Name, &t.Status)

		if err != nil {
			return nil, fmt.Errorf("error scannig rows %v", err)
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %v", err)
	}

	return tasks, nil
}

func DeleteTask(id int) error {
	_, err := storage.GetDB().Exec("DELETE FROM tasks WHERE ID = ?", id)

	return err
}

func PrintTasks(tasks []Task) {
	rows := struct2Strings(tasks)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("ID", "Name", "Status").
		Rows(rows...)

	fmt.Println(t)
}

func struct2Strings(tasks []Task) [][]string {
	var result [][]string

	for _, task := range tasks {
		status := fmt.Sprintf("%v", task.Status)
		id := fmt.Sprint(task.ID)
		row := []string{
			id,
			task.Name,
			status,
		}

		result = append(result, row)
	}
	return result
}

func ChangeStatus(id int, newStatus string) error {
	_, err := storage.GetDB().Exec("UPDATE tasks SET status = ? WHERE id = ?", newStatus, id)
	return err
}
