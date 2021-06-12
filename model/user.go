package model

type User struct {
	ID       int
	Username string
}

type Task struct {
	ID     int
	UserID int
	Name   string
}
