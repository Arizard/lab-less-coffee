package infra

import (
	"database/sql"
	"time"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository() {
	repo := MySQLUserRepository{}
	db, err := sql.Open("mysql", "root:password@localhost/quarry_bb")
	if err != nil {
		panic(err)
	}

	repo.db = db

	if err := repo.db.Ping(); err != nil {
		panic(err)
	}

	repo.db.SetConnMaxLifetime(time.Minute * 3)
	repo.db.SetMaxOpenConns(10)
	repo.db.SetMaxIdleConns(10)
}

func (repo *MySQLUserRepository) Get(uid entity.UserUID) (user *entity.User, err error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MySQLUserRepository) GetByLogin(login string) (user *entity.User, err error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MySQLUserRepository) Create(user entity.User) (createdUser *entity.User, err error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MySQLUserRepository) Delete(uid entity.UserUID) (err error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MySQLUserRepository) Update(user entity.User) (updatedUser *entity.User, err error) {
	panic("not implemented") // TODO: Implement
}
