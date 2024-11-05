package route

import (
	"flag"
	"strconv"
	"task-tracker-cli/internal/service"
)

func Route() {
	flag.Parse()
	args := flag.Args()
	if args[0] == "add" {
		service.Add(args[1])
	}
	if args[0] == "update" {
		id, _ := strconv.Atoi(args[1])
		service.Update(args[2], id)
	}
	if args[0] == "delete" {
		id, _ := strconv.Atoi(args[1])
		service.Delete(id)
	}
	if args[0] == "mark-in-progress" {
		id, _ := strconv.Atoi(args[1])
		service.Mark_In_Progress(id)
	}
	if args[0] == "mark-done" {
		id, _ := strconv.Atoi(args[1])
		service.Mark_Done(id)
	}
	if args[0] == "list" && len(args) == 1 {
		service.List()
		return
	}
	if args[0] == "list" && args[1] == "done" {
		service.ListDone()
	}
	if args[0] == "list" && args[1] == "in-progress" {
		service.ListInProgress()
	}
	if args[0] == "list" && args[1] == "todo" {
		service.ListTodo()
	}
}
