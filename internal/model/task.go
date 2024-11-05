package model

import "time"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

func GenEntity(id int, s string) Task {
	data := Task{
		Id:          id,
		Description: s,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return data
}
