package creational

import "fmt"

type User struct {
	name  string
	email string
}

type UserBuilder struct {
	user User
}

func (u *UserBuilder) SetEmail(email string) *UserBuilder {
	u.user.email = email
	return u
}

func (u *UserBuilder) SetName(name string) *UserBuilder {
	u.user.name = name
	return u
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) Build() User {
	return u.user
}

func builder() {
	user := NewUserBuilder().
		SetEmail("myEMail").
		SetName("myName").
		Build()

	fmt.Println(user)
}
