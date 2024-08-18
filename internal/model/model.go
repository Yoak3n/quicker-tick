package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Task struct {
		ID          string   `json:"id"`
		Description string   `json:"description"`
		Children    []string `json:"children"`
		Checked     bool     `json:"checked"`
		Action      string   `json:"action"`
		Date        string   `json:"date"`
		Root        bool     `json:"root"`
	}
	TasksTable struct {
		ID          string `gorm:"primarykey"`
		Description string
		Children    string
		Checked     bool
		Action      string
		Root        bool
		Date        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}

	TaskView struct {
		ID          string     `json:"id"`
		Description string     `json:"description"`
		Checked     bool       `json:"checked"`
		Children    []TaskView `json:"children"`
		Action      string     `json:"action"`
		Root        bool       `json:"root"`
	}
	Item struct {
		Type    string            `json:"type"`
		Content []Item            `json:"content,omitempty"`
		Text    string            `json:"text,omitempty"`
		Attrs   map[string]string `json:"attrs,omitempty"`
	}
)
