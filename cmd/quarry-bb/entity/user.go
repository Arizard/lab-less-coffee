package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
)

type UserUID string

type User struct {
	FullName     string
	ShortName    string
	UID          UserUID
	PasswordHash string
	Login        string
	Verified     bool
	Created      time.Time
}

type UserRegistration struct {
	Login    string
	Password string
}

func NewUserUID() UserUID {
	userUID := uuid.New()

	return UserUID(userUID.String())
}

func NewHashedPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func ComparePassword(password string, passwordHash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, passwordHash)
	if err != nil {
		return false
	}
	return match
}

type UserRepository interface {
	Get(uid UserUID) (user *User, err error)
	GetByLogin(login string) (user *User, err error)
	Create(user User) (createdUser *User, err error)
	Delete(uid UserUID) (err error)
	Update(user User) (updatedUser *User, err error)
}

type UserService struct {
	repo UserRepository
}

func (svc *UserService) UserRepository() (repo UserRepository, err error) {
	if svc.repo == nil {
		return nil, errors.New("repo is nil")
	}
	return svc.repo, nil
}

func (svc *UserService) Register(reg UserRegistration) (createdUser *User, err error) {
	existingUser, err := svc.repo.GetByLogin(reg.Login)
	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf("user with login %s already exists", reg.Login))
	}

	hash, err := NewHashedPassword(reg.Password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not hash password"))
	}

	user := User{
		UID:          NewUserUID(),
		Login:        reg.Login,
		PasswordHash: hash,
		FullName:     "New User",
		ShortName:    "New User",
		Created:      time.Now(),
	}

	createdUser, err = svc.repo.Create(user)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not register new user: %s", err))
	}

	return createdUser, nil
}

func (svc *UserService) NewSession(login string, password string) (session *UserSession, err error) {
	user, err := svc.repo.GetByLogin(login)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not create new session: %s", err))
	}
	if !ComparePassword(password, user.PasswordHash) {
		return nil, errors.New("could not log in")
	}
	return &UserSession{
		User:        user.UID,
		AccessToken: []byte("magic"),
		Created:     time.Now(),
	}, nil
}
