package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository() *MySQLUserRepository {
	repo := MySQLUserRepository{}
	db, err := sql.Open("mysql", "root:password@/quarry_bb")
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

	return &repo
}

type userRow struct {
	Id           int
	FullName     string
	ShortName    string
	Uid          string
	PasswordHash string
	Login        string
	Verified     bool
	Created      time.Time
}

func getByColumn(db *sql.DB, col string, val string) (*entity.User, error) {

	var user entity.User
	var created int64
	query := fmt.Sprintf("SELECT full_name, short_name, uid, password_hash, login, verified, created FROM user WHERE %s = '%s'", col, val)
	fmt.Println(query)
	row := db.QueryRow(query)
	err := row.Scan(
		&user.FullName,
		&user.ShortName,
		&user.UID,
		&user.PasswordHash,
		&user.Login,
		&user.Verified,
		&created,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrorNotFound
		}
		return nil, fmt.Errorf("could not get user with column '%s' matching '%s': %w", col, val, err)
	}

	user.Created = time.Unix(created, 0)

	return &user, nil
}

func (repo *MySQLUserRepository) Get(uid entity.UserUID) (*entity.User, error) {
	return getByColumn(repo.db, "uid", string(uid))
}

func (repo *MySQLUserRepository) GetByLogin(login string) (user *entity.User, err error) {
	return getByColumn(repo.db, "login", login)
}

func (repo *MySQLUserRepository) Create(user entity.User) (createdUser *entity.User, err error) {
	db := repo.db

	stmt, err := db.Prepare("INSERT INTO user(full_name, short_name, uid, password_hash, login, verified, created) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("could not prepare create statement: %w", err)
	}

	_, err = stmt.Exec(
		user.FullName,
		user.ShortName,
		user.UID,
		user.PasswordHash,
		user.Login,
		user.Verified,
		user.Created.Unix(),
	)

	if err != nil {
		return nil, fmt.Errorf("could not execute create statement: %w", err)
	}

	createdUser = &user

	return createdUser, nil
}

func (repo *MySQLUserRepository) Delete(uid entity.UserUID) (err error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MySQLUserRepository) Update(user entity.User) (updatedUser *entity.User, err error) {
	panic("not implemented") // TODO: Implement
}
