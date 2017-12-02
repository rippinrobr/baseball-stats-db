package databank

import (
	"errors"
	"os"
)

const TestFlagError = "error"

// InsertFuncForTesting is here to help test For
func InsertFuncForTesting(flag string, d interface{}) error {
	if flag == "error" {
		return errors.New("Throwing an error as requested")
	}

	return nil
}

func ParserTestingFunc(dummy *os.File, flag interface{}) error {
	if flag.(string) == "error" {
		return errors.New("Throwing an error as requested")
	}

	return nil
}

type RepositoryMock struct{}

func (r *RepositoryMock) CloseConn() error {
	return nil
}

func (r *RepositoryMock) Insert(flag string, d interface{}) error {
	if flag == "error" {
		return errors.New("returning an error as requested")
	}

	return nil
}

func (r *RepositoryMock) Truncate(flag string) error {
	if flag == "error" {
		return errors.New("returning an error as requested")
	}

	return nil
}
