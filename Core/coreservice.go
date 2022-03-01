package core

import (
	"github.com/wathuta/myservice"
)

type Service struct {
	//postgres db connection  a gorm
	m map[int64]myservice.User
}

func NewService() myservice.Service {
	return &Service{
		m: map[int64]myservice.User{
			1: {UserID: 1, Name: "alice"},
		},
	}
}
func (s *Service) GetUser(id int64) (myservice.User, error) {
	//should query the database
	var user myservice.User

	if result, ok := s.m[id]; ok {
		return result, nil
	}
	return user, myservice.ErrNotFound
}
func (s *Service) GetUsers(ids []int64) (map[int64]myservice.User, error) {
	//should query the db
	result := map[int64]myservice.User{}

	for _, id := range ids {
		if u, ok := s.m[id]; ok {
			result[id] = s.m[u.UserID]
		}
	}
	return result, myservice.ErrNotFound
}
