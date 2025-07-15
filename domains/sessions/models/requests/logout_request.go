package requests

type LogoutRequest struct {
	IsRevoked string `json:"is_revoked" validate:"required"`
}
