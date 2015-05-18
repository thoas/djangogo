package models

import (
	"time"
)

type User struct {
	ID          int64
	Username    string    `sql:"not null; unique; type:varchar(80)"`
	FirstName   string    `sql:"not null; type:varchar(50)"`
	LastName    string    `sql:"not null; type:varchar(50)"`
	Email       string    `sql:"not null; type:varchar(75)"`
	Password    string    `sql:"not null; type:varchar(128)"`
	IsStaff     bool      `sql:"not null"`
	IsActive    bool      `sql:"not null"`
	IsSuperuser bool      `sql:"not null"`
	LastLogin   time.Time `sql:"not null"`
	DateJoined  time.Time `sql:"not null"`
}

func (u *User) TableName() string {
	return "auth_user"
}

func (u *User) IsAuthenticated() bool {
	return u.ID > 0
}
