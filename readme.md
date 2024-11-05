## Việt Nam version

### Tóm tắt

- Đây là dự án cho bài tập backend trên [https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker)
- Dự án không sử dụng thư viện ngoài

### Những khó khăn khi làm dự án, những gì học được

- Tạo model

```go
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
```

- Đọc ghi file

```go
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
```
