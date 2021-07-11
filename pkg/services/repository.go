package services

import "errors"

type RepositoryError error

var (
	RepositoryRecordNotFound RepositoryError = errors.New("record not found")
)
