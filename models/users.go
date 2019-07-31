package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/ik5/go-into/types"
)

// User data structure
type User struct {
	ID          uint64         `json:"id" db:"id"`
	Roles       types.Role     `json:"roles" db:"roles"`
	Username    string         `json:"username" db:"username"`
	Password    string         `json:"password" db:"password"`
	Email       string         `json:"email" db:"email"`
	Name        sql.NullString `json:"name" db:"name"`
	IconAddress sql.NullString `json:"icon_address,omitempty" db:"icon_address"`
	Enabled     bool           `json:"-" db:"enabled"`
	Deleted     bool           `json:"-" db:"deleted"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// NullUser is a representation for a user struct that can be null at db level
type NullUser struct {
	User  User
	Valid bool
}

// Users holds a list of current users
type Users []NullUser

// String implement interface lookup for String to display data type as string
func (u User) String() string {
	if u.Name.Valid {
		return fmt.Sprintf("%d - %s (%s)", u.ID, u.Name.String, u.Username)
	}
	return fmt.Sprintf("%d - %s (%s)", u.ID, u.Email, u.Username)
}

// Scan implements the Scanner interface.
func (nu *NullUser) Scan(value interface{}) error {
	if value == nil {
		nu.User, nu.Valid = User{}, false
		return nil
	}
	nu.Valid = true
	nu.User = value.(User)
	return nil
}

// Value implements the driver Valuer interface.
func (nu *NullUser) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	return nu.User, nil
}
