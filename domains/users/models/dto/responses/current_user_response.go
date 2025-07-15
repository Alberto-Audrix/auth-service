package responses

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type CurrentUserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
