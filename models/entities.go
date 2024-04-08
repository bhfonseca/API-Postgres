package models

type Todo struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}
