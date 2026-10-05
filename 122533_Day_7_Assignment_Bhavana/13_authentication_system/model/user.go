package model

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
}
