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
	Password string `json:"password"`
}

type LoginRequest struct {
	Email     string `json:"email"`
	Passsword string `json:"password"`
}

type GetByEmailResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
