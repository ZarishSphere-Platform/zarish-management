package models

import "time"

// Project represents a project
type Project struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	Code        string `gorm:"size:50;uniqueIndex" json:"code"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"size:50;default:'planning'" json:"status"` // planning, active, on_hold, completed, cancelled
	Priority    string `gorm:"size:50;default:'medium'" json:"priority"`

	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`

	ManagerID uint `gorm:"index" json:"manager_id"`
}

// Task represents a project task
type Task struct {
	BaseModel

	ProjectID   uint   `gorm:"index;not null" json:"project_id"`
	Name        string `gorm:"size:500;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"size:50;default:'todo'" json:"status"` // todo, in_progress, review, done
	Priority    string `gorm:"size:50;default:'medium'" json:"priority"`

	AssigneeID *uint      `gorm:"index" json:"assignee_id,omitempty"`
	DueDate    *time.Time `json:"due_date,omitempty"`

	EstimatedHours float64 `json:"estimated_hours"`
	ActualHours    float64 `json:"actual_hours"`
}

// TableName overrides
func (Project) TableName() string {
	return "projects"
}

func (Task) TableName() string {
	return "tasks"
}
