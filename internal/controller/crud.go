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
	db.Where("date IN ?", date).Find(&tasks)
	return tasks

}
