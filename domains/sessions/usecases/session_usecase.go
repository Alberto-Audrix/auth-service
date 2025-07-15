package usecases

import "bootcamp-content-interaction-service/domains/sessions"

type sessionUseCase struct {
	repo sessions.SessionRepository
}

func NewSessionUseCase(repo sessions.SessionRepository) sessions.SessionUseCase {
	return sessionUseCase{repo: repo}
}
