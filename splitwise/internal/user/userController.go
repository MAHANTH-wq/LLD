package user

import (
	"fmt"
	"splitwise/internal/utils"
)

type UserController struct {
	users []*User
}

func NewUserController() *UserController {
	return &UserController{
		users: make([]*User, 0),
	}
}

func (uc *UserController) ListAllUsers() {
	for _, user := range uc.users {

		fmt.Println("User Id: ", user.id, " User Name: ", user.name)
	}
}

func (uc *UserController) AddNewUser(uName string) *User {

	id := len(uc.users) + 1

	newUser := NewUser(id, uName)
	uc.users = append(uc.users, newUser)
	return newUser
}

func (uc *UserController) DeleteUser(u *User) {

	for index, user := range uc.users {

		if user.GetUserName() == u.GetUserName() && user.GetUserId() == u.GetUserId() {
			uc.users = append(uc.users[0:index], uc.users[index+1:]...)
			return
		}
	}

	fmt.Println("There is no user with name: ", u.GetUserName())
}

func (uc *UserController) GetUser(userId int) *User {

	for _, user := range uc.users {
		if user.id == userId {
			return user
		}
	}
	fmt.Println("No user with id ", userId)
	return nil
}

func (uc *UserController) GetNewIndividualSplitInput() (*User, *User, int, error) {

	fmt.Println("Select the paid user id from below")
	uc.ListAllUsers()
	fromUserId, err := utils.GetIntegerInput()

	if err != nil {
		return nil, nil, 0, err
	}
	fmt.Println("Select the reciever user id from below")
	uc.ListAllUsers()
	toUserId, err := utils.GetIntegerInput()

	if err != nil {
		return nil, nil, 0, err
	}

	fmt.Println("Enter the amount paid")
	amount, err := utils.GetIntegerInput()

	if err != nil {
		return nil, nil, 0, err
	}

	return uc.GetUser(fromUserId), uc.GetUser(toUserId), amount, nil
}
