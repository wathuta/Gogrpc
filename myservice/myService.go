package myservice

import "errors"

var ErrNotFound = errors.New("not found")

type User struct {
	UserID int64
	Name   string
}

type Service interface {
	GetUser(id int64) (User, error)
	GetUsers(ids []int64) (map[int64]User, error)
}

// func NewService(UserID int64, Name string) Service {
// 	return &User{
// 		UserID: UserID,
// 		Name:   Name,
// 	}
// }
// func (u *User) GetUser(id int64) (User, error) {
// 	var user User
// 	return user, nil
// }
// func (u *User) GetUsers(ids []int64) (map[int64]User, error) {
// 	users := make(map[int64]User)

// 	return users, nil
// }
