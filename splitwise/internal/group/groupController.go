package group

import (
	"fmt"
	"splitwise/internal/expense"
	"splitwise/internal/split"
	"splitwise/internal/user"
)

type GroupController struct {
	groups []*Group
}

func NewGroupController() *GroupController {
	return &GroupController{
		groups: make([]*Group, 0),
	}
}

func (gc *GroupController) CheckGroupIdExists(groupId int) bool {

	for _, group := range gc.groups {

		if group.id == groupId {
			return true
		}
	}

	fmt.Println("Group Id does not exist")
	return false

}

func (gc *GroupController) AddExpenseToGroup(groupId int, e *expense.Expense) {

	for _, group := range gc.groups {
		if group.id == groupId {
			group.addExpenseToGroup(e)
			return
		}
	}
}

func (gc *GroupController) ListAllGroups() {

	for _, group := range gc.groups {
		fmt.Println("Group Id: ", group.id, "Group Name: ", group.name)
	}
}

func (gc *GroupController) AddNewGroup(gName string) {

	id := len(gc.groups) + 1
	newGroup := createNewGroup(id, gName)
	gc.groups = append(gc.groups, newGroup)
}

func (gc *GroupController) ListUsersFromGroup(groupId int) {

	for _, group := range gc.groups {
		if group.id == groupId {
			fmt.Println("The following are the users in this group ", group.name)
			group.listAllUsers()
			return
		}
	}
	fmt.Println("No Group with Id ", groupId)
	return
}

func (gc *GroupController) AddSplitsToUsersFromGroup(groupId int) ([]*split.Split, error) {

	for _, group := range gc.groups {
		if group.id == groupId {
			fmt.Println("Add Split Amount to the users of the group ", group.name)
			splitList, err := group.addSplitToUsers()
			return splitList, err
		}
	}
	return nil, fmt.Errorf("No Group with Id %d", groupId)
}

func (gc *GroupController) AddUserToGroup(id int, u *user.User) {

	for _, group := range gc.groups {

		if group.id == id {
			group.addUserToGroup(u)
			return
		}
	}

	fmt.Println("There is no group with id ", id)
}

func (gc *GroupController) RemoveUserToGroup(id int, u *user.User) {

	for _, group := range gc.groups {

		if group.id == id {
			group.removeUserFromGroup(u)
			return
		}
	}

	fmt.Println("There is no group with id ", id)
	return
}

func (gc *GroupController) DeleteGroup(id int) {

	for index, group := range gc.groups {

		if group.id == id {
			gc.groups = append(gc.groups[0:index], gc.groups[index+1:]...)
			fmt.Println("Deleted Group with id and name: ", id, group.name)
			return
		}

	}

	fmt.Println("There is no group with id ", id)
	return
}
