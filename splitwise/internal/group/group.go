package group

import (
	"fmt"
	"splitwise/internal/expense"
	"splitwise/internal/split"
	"splitwise/internal/user"
	"splitwise/internal/utils"
)

type Group struct {
	id               int
	name             string
	users            []*user.User
	groupExpenseList []*expense.Expense
}

func createNewGroup(id int, name string) *Group {
	return &Group{
		id:               id,
		name:             name,
		users:            make([]*user.User, 0),
		groupExpenseList: make([]*expense.Expense, 0),
	}
}

func (g *Group) addExpenseToGroup(e *expense.Expense) {
	g.groupExpenseList = append(g.groupExpenseList, e)
}

func (g *Group) listAllUsers() {

	for _, user := range g.users {
		fmt.Println("User Id: ", user.GetUserId(), " User Name: ", user.GetUserName())
	}
	return
}

func (g *Group) addSplitToUsers() ([]*split.Split, error) {

	listOfSplits := make([]*split.Split, 0)
	for _, user := range g.users {
		fmt.Println("Enter the split amount for User Id: ", user.GetUserId(), " User Name: ", user.GetUserName())
		splitAmount, err := utils.GetIntegerInput()

		if err != nil {
			fmt.Println("Invalid Split Amount")
			return nil, err
		}

		if splitAmount == 0 {
			continue
		}
		split := split.CreateNewSplit(user, splitAmount)
		listOfSplits = append(listOfSplits, split)

	}
	return listOfSplits, nil
}

func (g *Group) addUserToGroup(u *user.User) {
	g.users = append(g.users, u)
}

func (g *Group) removeUserFromGroup(u *user.User) {

	for index, groupUser := range g.users {

		if groupUser.GetUserName() == u.GetUserName() && groupUser.GetUserId() == u.GetUserId() {
			g.users = append(g.users[0:index], g.users[index+1:]...)
			return
		}
	}

	fmt.Println("No User with name ", u.GetUserName(), " in the group")
	return
}
