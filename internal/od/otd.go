package od

import (
	"errors"
	"quicker-tick/internal/controller"
	"quicker-tick/internal/model"
	"strings"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

// 解析json
func UnmashalRawString(object string, date string) error {
	result := gjson.Parse(object)
	if result.Get("type").String() != "doc" {
		return errors.New("type is not doc")
	}

	list := result.Get("content").Array()
	if len(list) == 0 {
		return errors.New("content is empty")
	}
	tasks := unmashalTaskList(list[0], date)
	for _, v := range tasks {
		v.Root = true
		controller.UpdateTaskAsRoot(v.ID)
	}
	return nil
}

func unmashalTaskList(result gjson.Result, date string) []*model.Task {
	content := result.Get("content").Array()
	tasks := make([]*model.Task, 0)
	for _, v := range content {
		if t := v.Get("type").String(); t == "taskItem" {
			task := model.Task{
				ID:   uuid.NewString(),
				Date: date,
			}
			task.Checked = v.Get("attrs.checked").Bool()
			content1 := v.Get("content").Array()
			for _, v1 := range content1 {
				if t := v1.Get("type").String(); t == "paragraph" {
					task.Description = unmashalText(v1)
				} else if t == "taskList" {
					children := unmashalTaskList(v1, date)
					if len(children) > 0 {
						for _, child := range children {
							task.Children = append(task.Children, child.ID)
						}

					}
				}
			}
			tasks = append(tasks, &task)
			childern := strings.Join(task.Children, ",")
			taskRecord := &model.TasksTable{
				ID:          task.ID,
				Description: task.Description,
				Children:    childern,
				Checked:     task.Checked,
				Root:        task.Root,
				Action:      task.Action,
				Date:        task.Date,
			}
			controller.CreateTask(taskRecord)
		}
	}
	return tasks
}

func unmashalText(result gjson.Result) string {
	text := result.Get("content.0.text").String()
	return text
}
