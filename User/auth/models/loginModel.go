package models

type LoginData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}
