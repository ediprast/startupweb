package user

import "time"

type User struct {
	Id         int
	Name       string
	Occupation string
	Email      string
	Token      string
	Password   string
	Avatar     string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
