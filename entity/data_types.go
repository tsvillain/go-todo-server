package entity

// Enum for Priority
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

// Todo DataStructure
type Todo struct {
	Id       int      `json:"id"`
	Task     string   `json:"task"`
	Status   bool     `json:"status"`
	UserName   string   `json:"username"`
	Priority Priority `json:"priority"`
}
