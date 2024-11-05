package service

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker-cli/internal/model"
)

func Read() model.Tasks {
	byte, _ := os.ReadFile("./task.json")
	var data model.Tasks
	json.Unmarshal(byte, &data)
	return data
}

func Write(tasks model.Tasks) {
	bytes, _ := json.Marshal(tasks)
	os.WriteFile("./task.json", bytes, 0644)
}

func Add(s string) {
	data := Read()
	n := len(data.Tasks)
	data.Tasks = append(data.Tasks, model.GenEntity(n+1, s))
	Write(data)
}

func Update(s string, id int) {
	data := Read()
	n := len(data.Tasks)
	for i := 0; i < n; i++ {
		if data.Tasks[i].Id == id {
			data.Tasks[i].Description = s
		}
	}
	Write(data)
}

func Delete(id int) {
	data := Read()
	n := len(data.Tasks)

	newData := model.Tasks{}

	for i := 0; i < n; i++ {
		if data.Tasks[i].Id != id {
			newData.Tasks = append(newData.Tasks, data.Tasks[i])
		}
	}

	Write(newData)
}

func Mark_In_Progress(id int) {
	// status -> 1
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		if data.Tasks[i].Id == id {
			data.Tasks[i].Status = 1
		}
	}

	Write(data)
}

func Mark_Done(id int) {
	// status -> 2
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		if data.Tasks[i].Id == id {
			data.Tasks[i].Status = 2
		}
	}
}

func List() {
	g := map[int]string{
		0: "todo",
		1: "in progress",
		2: "completed",
	}
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		fmt.Println(data.Tasks[i].Id, data.Tasks[i].Description, g[data.Tasks[i].Status])
	}
}

func ListDone() {
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		if data.Tasks[i].Status == 2 {
			fmt.Println(data.Tasks[i].Id, data.Tasks[i].Description)
		}
	}
}

func ListTodo() {
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		if data.Tasks[i].Status == 0 {
			fmt.Println(data.Tasks[i].Id, data.Tasks[i].Description)
		}
	}
}

func ListInProgress() {
	data := Read()
	n := len(data.Tasks)

	for i := 0; i < n; i++ {
		if data.Tasks[i].Status == 1 {
			fmt.Println(data.Tasks[i].Id, data.Tasks[i].Description)
		}
	}
}
