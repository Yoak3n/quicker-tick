package controller

import (
	"fmt"
	"quicker-tick/internal/database"
	"quicker-tick/internal/model"
)

func CreateTask(task *model.TasksTable) {
	db := database.GetDB()
	db.Create(task)
}

func UpdateTaskAsRoot(id string) {
	db := database.GetDB()
	db.Model(&model.TasksTable{}).Where("id = ?", id).Update("root", true)
	fmt.Println("Updated task as root")
}

func ReadTasksByDate(date []string) []model.TasksTable {
	db := database.GetDB()
	var tasks []model.TasksTable
	db.Where("due_date IN ?", date).Find(&tasks)
	return tasks
}

func CreateAction(action *model.ActionsTable) {
	db := database.GetDB()
	db.Create(action)
}
func ReadActionByID(id string) *model.ActionsTable {
	db := database.GetDB()
	var action *model.ActionsTable
	db.Where("id = ?", id).First(&action)
	return action
}

func ReadAllActions() []*model.ActionsTable {
	db := database.GetDB()
	var actions []*model.ActionsTable
	db.Find(&actions)
	return actions
}
