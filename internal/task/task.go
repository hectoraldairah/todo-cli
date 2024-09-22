package task

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

func AddTask(name string, currentTasks []Task) []Task {

	task := Task{
		ID:     len(currentTasks) + 1,
		Name:   name,
		Status: ToDo,
	}

	currentTasks = append(currentTasks, task)

	fmt.Printf("\nTask \"%v\" was added\n", task.Name)

	return currentTasks
}

func PrintTasks(currentTasks []Task) {
	rows := struct2Strings(currentTasks)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("ID", "Name", "Status").
		Rows(rows...)

	fmt.Println(t)
}

func struct2Strings(currentTasks []Task) [][]string {
	var result [][]string

	for _, task := range currentTasks {
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

func RemoveTask(id string, currentTasks []Task) []Task {
	convertedId, _ := strconv.Atoi(id)
	for index, task := range currentTasks {
		if convertedId == task.ID {
			currentTasks = append(currentTasks[:index], currentTasks[index+1:]...)
		}
	}
	return currentTasks
}
