package user

type User struct {
	id   int
	name string
}

func NewUser(id int, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u *User) GetUserId() int {
	return u.id
}

func (u *User) GetUserName() string {
	return u.name
}
