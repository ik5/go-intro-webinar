package db

import (
	"database/sql"
	"errors"
	"testing"
)

func TestIsNoRowsReturned(t *testing.T) {
	err := sql.ErrNoRows

	if !IsNoRows(err) {
		t.Errorf("Gave err of %s but got false on IsNoRows", err)
	}
}

func TestIsNoRowsNotReturned(t *testing.T) {
	err := errors.New("Just an error")

	if IsNoRows(err) {
		t.Errorf("Gave err of %s but got true on IsNoRows", err)
	}
}

func TestIsNoRowsErrNil(t *testing.T) {
	var err error

	if IsNoRows(err) {
		t.Errorf("Gave err of %T but got true on IsNoRows", err)
	}
}
