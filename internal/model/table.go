package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	TasksTable struct {
		ID          string `gorm:"primarykey"`
		Title       string
		Description string
		Children    string
		Checked     bool
		Tags        string
		Priority    int
		Status      int
		Action      string
		Root        bool
		DueDate     string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}
	ActionsTable struct {
		ID        string `gorm:"primarykey"`
		Name      string
		Command   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	TagsTable struct {
		gorm.Model
		Title string
		Color string
	}
)
