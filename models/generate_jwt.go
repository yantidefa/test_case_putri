package models

type GenerateJWT struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
