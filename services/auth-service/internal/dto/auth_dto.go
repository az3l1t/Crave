package dto

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

type LoginRequest struct {
	Email     string `json:"username"`
	Passsword string `json:"passsword"`
}

type GetByEmailResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
