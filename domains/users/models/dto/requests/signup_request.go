package requests

type SignUpRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Bio             string `json:"bio"`
	Gender          string `json:"gender"`
	Phone           string `json:"phone"`
	Country         string `json:"country"`
	Profile         string `json:"profile"`
}
