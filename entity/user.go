package entity

import (
	"ostadbun/pkg/hash"
)

type User struct {
	//must be hash
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *User) Email_Hashe() string {
	return hash.Hasher(u.Email)
}
