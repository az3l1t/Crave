package dto

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Passsword string `json:"passsword"`
}

type LoginRequest struct {
	Username  string `json:"username"`
	Passsword string `json:"passsword"`
}

type GetByEmailResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
