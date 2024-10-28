package od

import (
	"quicker-tick/internal/controller"
	"quicker-tick/internal/model"
	"strconv"
	"strings"
)

func GetTasks(dates []string) map[string][]*model.TaskView {
	query := make(map[string]model.TasksTable, 0)
	view := make(map[string][]*model.TaskView)
	if len(dates) == 0 {
		return view
	}

	ts := controller.ReadTasksByDate(dates)
	for _, t := range ts {
		query[t.ID] = t
	}
	for _, t := range ts {
		if t.Root {
			taskView := &model.TaskView{
				ID:          t.ID,
				Description: t.Description,
				Checked:     t.Checked,
				// Action:      t.Action,
				Root: true,
			}
			ParseTasksChildren(t, &query, taskView)
			view[t.DueDate] = append(view[t.DueDate], taskView)
		}
	}
	return view
}

// 递归解析子任务
func ParseTasksChildren(task model.TasksTable, query *map[string]model.TasksTable, tt *model.TaskView) {
	childrenIDs := strings.Split(task.Children, ",")
	for _, childID := range childrenIDs {
		if child, ok := (*query)[childID]; ok {
			childView := model.TaskView{
				ID:          child.ID,
				Description: child.Description,
				Checked:     child.Checked,
				// Actions:      child.Actions,
			}
			if len(child.Children) > 0 {
				ParseTasksChildren(child, query, &childView)
			}
			tt.Children = append(tt.Children, &childView)
		}
	}
}

func ConvertToDoc(tasks []*model.TaskView) model.Item {
	item := model.Item{
		Type: "doc",
		Content: []model.Item{
			{
				Type:    "taskList",
				Content: []model.Item{},
			},
		},
	}
	if len(tasks) == 0 {
		return item
	} else {
		for _, task := range tasks {
			item.Content[0].Content = append(item.Content[0].Content, convertToTaskList(task))
		}
	}
	return item
}

func ConvertToHtmlString(item model.Item, flag bool) (htmlString string) {
	if item.Type == "doc" {
		for _, child := range item.Content {
			htmlString += ConvertToHtmlString(child, false)
		}
	} else if item.Type == "taskList" && len(item.Content) > 0 {
		htmlString = `<ul data-type="taskList">`
		for _, child := range item.Content {
			htmlString += ConvertToHtmlString(child, true)
		}
		htmlString += `</ul>`
	} else if item.Type == "taskItem" {
		htmlString = `<li data-type="taskItem" data-checked="` + strconv.FormatBool(item.Attrs["checked"].(bool)) + `" id="` + (item.Attrs["id"].(string)) + `">`
		for _, child := range item.Content {
			htmlString += ConvertToHtmlString(child, false)
		}
		htmlString += `</li>`
	} else if item.Type == "paragraph" && !flag {
		for _, child := range item.Content {
			htmlString += ConvertToHtmlString(child, false)
		}
	} else if item.Type == "text" {
		htmlString += item.Text
	}
	return htmlString
}

func convertToTaskList(task *model.TaskView) model.Item {
	item := model.Item{
		Type: "taskItem",
		Content: []model.Item{
			{
				Type: "paragraph",
				Content: []model.Item{
					{
						Type: "text",
						Text: task.Description,
					},
				},
			},
		},
		Attrs: map[string]any{
			"checked": task.Checked,
			"id":      task.ID,
		},
	}
	if len(task.Children) > 0 {
		item.Content = append(item.Content, model.Item{
			Type:    "taskList",
			Content: []model.Item{},
		})
		for _, child := range task.Children {
			item.Content[1].Content = append(item.Content[1].Content, convertToTaskList(child))
		}
	}
	return item
}
