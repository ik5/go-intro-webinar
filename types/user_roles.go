package types

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"strconv"
)

// Role is a data type for holding user roles
type Role uint64

// A list of user Roles
const (
	RoleNone        Role = 1 << 0
	RoleCreate           = 1 << 1
	RoleEdit             = 1 << 2
	RoleReview           = 1 << 3
	RoleDelete           = 1 << 4
	RolePublish          = 1 << 5
	RoleCreateUser       = 1 << 6
	RoleManageUser       = 1 << 7
	RoleDisableUser      = 1 << 8
	RoleDeleteUser       = 1 << 9
)

// A list of common roles
const (
	RoleEditor Role = RoleEdit | RoleReview | RolePublish
	RoleCRUD   Role = RoleEditor | RoleCreate | RoleDelete
	RoleAdmin  Role = RoleCreate | RoleManageUser | RoleDisableUser
	RoleRoot   Role = RoleCRUD | RoleAdmin | RoleDeleteUser
)

// Scan implements the Scanner interface
func (r *Role) Scan(value interface{}) error {
	if value == nil {
		*r = RoleNone
		return nil
	}

	switch value.(type) {
	case uint8, uint, uint16, uint32, uint64:
		val := reflect.ValueOf(value)
		*r = Role(val.Uint())
		return nil

	case int8, int, int16, int32, int64:
		val := reflect.ValueOf(value)
		i := val.Int()
		if i > 0 {
			*r = Role(uint64(i))
			return nil
		}
		return errors.New("Invalid value for role was provided")

	case string:
		n, err := strconv.Atoi(value.(string))
		if err != nil {
			return err
		}
		if n >= 0 {
			*r = Role(n)
			return nil
		}
		return errors.New("Invalid value was provided for role")
	}
	return errors.New("Unsupported data type")
}

// Value implements the driver Valuer interface.
func (r *Role) Value() (driver.Value, error) {
	return uint64(*r), nil
}
