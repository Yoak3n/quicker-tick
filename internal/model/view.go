package model

type TaskView struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      int         `json:"status"`
	Priority    int         `json:"priority"`
	DueDate     string      `json:"due_date"`
	Actions     []*Action   `json:"actions"`
	Children    []*TaskView `json:"children"`
	Tags        []string    `json:"tags"`
	Checked     bool        `json:"checked"`
	Root        bool        `json:"root"`
}

type Action struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Command     string `json:"command"`
	Type        string `json:"type"`
}

type Tag struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
