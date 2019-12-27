package model

type User struct {
	id       string
	password string
}

func NewUser(id string, password string) (*User, error) {
	return &User{id: id, password: password}, nil
}

func (user *User) Id() string {
	if user != nil {
		return user.id
	}
	return ""
}

func (user *User) Password() string {
	if user != nil {
		return user.password
	}
	return ""
}
