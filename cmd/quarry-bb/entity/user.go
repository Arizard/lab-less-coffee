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
	FullName     string    `json:"full_name"`
	ShortName    string    `json:"short_name"`
	UID          UserUID   `json:"uid"`
	PasswordHash string    `json:"-"`
	Login        string    `json:"login"`
	Verified     bool      `json:"verified"`
	Created      time.Time `json:"created"`
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

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) UserRepository() UserRepository {
	return svc.repo
}

func (svc *UserService) Register(reg UserRegistration) (createdUser *User, err error) {
	existingUser, err := svc.repo.GetByLogin(reg.Login)
	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf("user with login %s already exists", reg.Login))
	}

	if err != nil {
		if !errors.Is(err, ErrorNotFound) {
			return nil, fmt.Errorf("could not query existing employee: %w", err)
		}
	}

	hash, err := NewHashedPassword(reg.Password)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
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
		return nil, fmt.Errorf("could not register new user: %w", err)
	}

	return createdUser, nil
}

func (svc *UserService) NewSession(login string, password string) (session *UserSession, err error) {
	user, err := svc.repo.GetByLogin(login)
	if err != nil {
		return nil, fmt.Errorf("could not create new session: %s", err)
	}
	if !ComparePassword(password, user.PasswordHash) {
		return nil, errors.New("could not find user with those credentials")
	}
	return &UserSession{
		User:        user.UID,
		AccessToken: "magic",
		Created:     time.Now(),
	}, nil
}
