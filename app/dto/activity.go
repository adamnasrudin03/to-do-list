package dto

// Struct req create activity
type CreateActivity struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// Struct req Update activity
type UpdateActivity struct {
	Title string `json:"title"`
	Email string `json:"email" validate:"email"`
}
