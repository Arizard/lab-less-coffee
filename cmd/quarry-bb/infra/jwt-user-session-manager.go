package infra

import "errors"

type JWTUserSessionManager struct{}

func NewJWTUserSessionManager() *JWTUserSessionManager {
	return &JWTUserSessionManager{}
}

func (m *JWTUserSessionManager) NewAccessToken(login string) string {
	return "magic"
}

func (m *JWTUserSessionManager) ValidateAccessToken(token string, errCb func(err error)) (isValid bool) {
	if token == "magic" {
		return true
	}

	errCb(errors.New("access token was incorrect"))
	return false
}
