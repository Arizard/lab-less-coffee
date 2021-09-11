package entity

import (
	"time"
)

type UserSession struct {
	User        UserUID
	AccessToken []byte
	Created     time.Time
}
