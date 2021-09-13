package infra

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
)

type JSONUserRepository struct {
	dbPath string
	db     []entity.User
}

func NewJSONUserRepository(dbPath string) *JSONUserRepository {
	repo := JSONUserRepository{
		dbPath: dbPath,
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := ioutil.WriteFile(dbPath, []byte("[]"), os.FileMode(0644))
		if err != nil {
			panic(err)
		}
	}

	contents, err := ioutil.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &repo.db); err != nil {
		panic(err)
	}

	return &repo
}

func (repo *JSONUserRepository) Save() {
	contents, err := json.Marshal(repo.db)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(repo.dbPath, contents, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

func (repo *JSONUserRepository) Get(uid entity.UserUID) (user *entity.User, err error) {
	for _, user := range repo.db {
		if user.UID == uid {
			return &user, nil
		}
	}

	return nil, errors.New("not found")
}

func (repo *JSONUserRepository) GetByLogin(login string) (user *entity.User, err error) {
	for _, user := range repo.db {
		if user.Login == login {
			return &user, nil
		}
	}

	return nil, errors.New("not found")
}

func (repo *JSONUserRepository) Create(user entity.User) (createdUser *entity.User, err error) {
	repo.db = append(repo.db, user)
	createdUser, _ = repo.Get(user.UID)
	repo.Save()
	return createdUser, nil
}

func (repo *JSONUserRepository) Delete(uid entity.UserUID) (err error) {
	userIndex := -1

	for index, existingUser := range repo.db {
		if existingUser.UID == uid {
			userIndex = index
		}
	}

	if userIndex != -1 {
		copy(repo.db[userIndex:], repo.db[userIndex+1:])
		repo.db[len(repo.db)-1] = entity.User{}
		repo.db = repo.db[:len(repo.db)-1]
		repo.Save()
		return nil
	}

	return errors.New("could not delete")

}

func (repo *JSONUserRepository) Update(user entity.User) (updatedUser *entity.User, err error) {
	userIndex := -1

	for index, existingUser := range repo.db {
		if existingUser.UID == user.UID {
			userIndex = index
		}
	}

	if userIndex != -1 {
		repo.db[userIndex] = user
		updatedUser, _ := repo.Get(user.UID)
		repo.Save()
		return updatedUser, nil
	}

	return &user, errors.New("not found")
}
