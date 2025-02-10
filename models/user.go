package models

type UserResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"upadated_at"`
}

type UserRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}