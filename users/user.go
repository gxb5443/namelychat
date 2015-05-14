package users

import (
	"strings"

	"github.com/satori/go.uuid"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Alias     string
	//Rooms map[hub]bool
}

func New(fname, lname string) *User {
	a := strings.Title(fname) + strings.Title(lname)
	return &User{
		Id:        uuid.NewV4().String(),
		FirstName: fname,
		LastName:  lname,
		Alias:     a,
	}
}

func aliasExists(a string) bool {
	return false
}

func (u *User) Save() error {
	return nil
}

func Delete() error {
	return nil
}

func GetAll() ([]*User, error) {
	return nil, nil
}

func GetById(id string) (*User, error) {
	return nil, nil
}
