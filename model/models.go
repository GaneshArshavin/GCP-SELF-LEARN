package model

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type NullString struct {
	sql.NullString
}

//Scan implements the sql.Scanner interface.
func (v *NullString) Scan(src interface{}) error {
	return v.NullString.Scan(src)
}

//Value implements the driver.Valuer interface.
func (v NullString) Value() (driver.Value, error) {
	return v.NullString.Value()
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (v *NullTime) Scan(src interface{}) error {
	v.Time, v.Valid = src.(time.Time)
	if v.Valid {
		v.Time = v.Time.In(time.UTC).Truncate(time.Millisecond)
	}
	return nil
}

//Value implements the driver.Valuer interface.
func (v NullTime) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Time, nil
}

type NullBool struct {
	sql.NullBool
}

func (v *NullBool) Scan(src interface{}) error {
	return v.NullBool.Scan(src)
}

//Value implements the driver.Valuer interface.
func (v NullBool) Value() (driver.Value, error) {
	return v.NullBool.Value()
}

type AccountsUser struct {
	ID        NullString `json:"id"`
	Username  NullString `json:"username	"`
	Passowrd  NullString `json:"passowrd"`
	Email     NullString `json:"email"`
	CreatedAt NullTime   `json:"created_at"`
	UpdatedAt NullTime   `json:"updated_at"`
}
