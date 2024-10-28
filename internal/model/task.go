package model

type Item struct {
	Id      string         `json:"id,omitempty"`
	Type    string         `json:"type"`
	Content []Item         `json:"content,omitempty"`
	Text    string         `json:"text,omitempty"`
	Attrs   map[string]any `json:"attrs,omitempty"`
}

type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Tags        []string `json:"tags"`
	Priority    int      `json:"priority"`
	Children    []string `json:"children"`
	Checked     bool     `json:"checked"`
	Action      string   `json:"action"`
	DueDate     string   `json:"due_date"`
	Root        bool     `json:"root"`
}
