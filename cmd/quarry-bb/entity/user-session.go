package entity

import (
	"errors"
	"time"
)

type UserSession struct {
	User        UserUID
	AccessToken string
	Created     time.Time
}

var (
	ErrorInvalidAccessToken = errors.New("invalid access token")
	ErrorUserNotPermitted   = errors.New("not permitted")
)

// UserSessionManager has the combined responsibility of authn and authz
type UserSessionManager interface {
	NewAccessToken(login string) string
	ValidateAccessToken(token string, errCb func(err error)) (isValid bool)
}
