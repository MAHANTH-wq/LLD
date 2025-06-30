package main

import (
	"fmt"
	"splitwise/internal/expense"
	"splitwise/internal/group"
	"splitwise/internal/split"
	"splitwise/internal/user"
	"splitwise/internal/utils"
)

type expenseType int

const (
	groupExpense expenseType = iota
	nonGroupExpense
)

type splitwise struct {
	groupController   *group.GroupController
	userController    *user.UserController
	expenseController *expense.ExpenseController
}

func newSplitwiseApp() *splitwise {

	return &splitwise{
		groupController:   group.NewGroupController(),
		userController:    user.NewUserController(),
		expenseController: expense.NewExpenseController(),
	}
}

func (app *splitwise) addNewUser() {
	fmt.Println("Enter username")
	userName, _ := utils.GetStringInput()
	newUser := app.userController.AddNewUser(userName)
	app.expenseController.AddNewUser(newUser)
}

func (app *splitwise) addNewGroup() {
	fmt.Println("Enter groupName")
	groupName, _ := utils.GetStringInput()
	app.groupController.AddNewGroup(groupName)
}

func (app *splitwise) addUserToGroup() {
	fmt.Println("Select the groupid from below")
	app.groupController.ListAllGroups()

	groupId, _ := utils.GetIntegerInput()

	fmt.Println("Select the userId from below")

	app.userController.ListAllUsers()

	userId, _ := utils.GetIntegerInput()

	userToBeAdded := app.userController.GetUser(userId)

	app.groupController.AddUserToGroup(groupId, userToBeAdded)

}

func (app *splitwise) printBalanceSheet() {

	app.expenseController.PrintBalanceSheet()
}

func (app *splitwise) addExpense() error {

	fmt.Println("Enter the option\n 0 for group expense \n 1 for non group expense")

	option, ok := utils.GetIntegerInput()

	if ok != nil {
		fmt.Println("Try again by selecting Valid Option")
		return ok
	}

	switch expenseType(option) {
	case groupExpense:
		app.createGroupExpense()
	case nonGroupExpense:
		app.createNonGroupExpense()
	default:
		return fmt.Errorf("Try again by selecting valid option")
	}

	return nil

}

func (app *splitwise) createGroupExpense() error {
	fmt.Println("Select the group id from the following: ")
	app.groupController.ListAllGroups()
	groupId, ok := utils.GetIntegerInput()

	if ok != nil {
		return ok
	}

	exists := app.groupController.CheckGroupIdExists(groupId)

	if !exists {
		return fmt.Errorf("Invalid Group Id Entered")
	}

	app.groupController.ListUsersFromGroup(groupId)

	fmt.Println("Enter the user id of the paid user ")

	paidUserId, err := utils.GetIntegerInput()

	if err != nil {
		return err
	}

	paidUser := app.userController.GetUser(paidUserId)

	fmt.Println("Enter the total Amount to be split ")

	totalAmountPaid, err := utils.GetIntegerInput()

	if err != nil {
		return err
	}

	splitAmountsList, err := app.groupController.AddSplitsToUsersFromGroup(groupId)

	if err != nil {
		return err
	}

	expense := app.expenseController.AddNewExpense("Expense Name bla bla", paidUser, splitAmountsList, totalAmountPaid, split.UnequalSplit)

	app.groupController.AddExpenseToGroup(groupId, expense)

	return nil
}

func (app *splitwise) createNonGroupExpense() error {
	return nil
}
