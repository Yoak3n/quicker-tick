package od

import (
	"log"
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
		taskView := &model.TaskView{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Checked:     t.Checked,
			DueDate:     t.DueDate,
			Status:      t.Status,
			Priority:    t.Priority,
			Actions:     make([]*model.Action, 0),
			// Action:      t.Action,
			Root: true,
		}
		if t.Action != "" {
			actionRecord := controller.ReadActionByID(t.Action)
			actionView := &model.Action{
				Type:        actionRecord.Type,
				ID:          actionRecord.ID,
				Name:        actionRecord.Name,
				Description: actionRecord.Description,
				Icon:        actionRecord.Icon,
				Command:     actionRecord.Command,
			}
			log.Println(actionView)
			taskView.Actions = append(taskView.Actions, actionView)
		}
		view[t.DueDate] = append(view[t.DueDate], taskView)
		// if t.Root {
		if t.Tags != "" {
			tags := strings.Split(t.Tags, ",")
			taskView.Tags = tags
		} else {
			taskView.Tags = []string{}
		}
	}
	return view
}

func GetActions() []*model.Action {
	views := make([]*model.Action, 0)
	actions := controller.ReadAllActions()
	for _, v := range actions {
		view := &model.Action{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Icon:        v.Icon,
			Command:     v.Command,
			Type:        v.Type,
		}
		views = append(views, view)
	}
	return views
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
