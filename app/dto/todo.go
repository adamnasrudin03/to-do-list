package dto

type CreateTodo struct {
	ActivityGroupID uint64 ` json:"activity_group_id" validate:"required,min=1"`
	Title           string ` json:"title" validate:"required"`
	IsActive        bool   `json:"is_active" validate:"required"`
	Priority        string `json:"priority" validate:"omitempty"`
}

type UpdateTodo struct {
	ActivityGroupID uint64 ` json:"activity_group_id" validate:"omitempty,min=1"`
	Title           string ` json:"title" validate:"omitempty"`
	IsActive        bool   `json:"is_active" validate:"omitempty"`
	Priority        string `json:"priority" validate:"omitempty"`
}
