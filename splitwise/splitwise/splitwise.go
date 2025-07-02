package splitwise

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
)

type splitwise struct {
	groupController   *group.GroupController
	UserController    *user.UserController
	ExpenseController *expense.ExpenseController
}

func NewSplitwiseApp() *splitwise {

	return &splitwise{
		groupController:   group.NewGroupController(),
		UserController:    user.NewUserController(),
		ExpenseController: expense.NewExpenseController(),
	}
}

func (app *splitwise) AddNewUser() {
	fmt.Println("Enter username")
	userName, _ := utils.GetStringInput()
	newUser := app.UserController.AddNewUser(userName)
	app.ExpenseController.AddNewUser(newUser)
}

func (app *splitwise) AddNewGroup() {
	fmt.Println("Enter groupName")
	groupName, _ := utils.GetStringInput()
	app.groupController.AddNewGroup(groupName)
}

func (app *splitwise) AddUserToGroup() {
	fmt.Println("Select the groupid from below")
	app.groupController.ListAllGroups()

	groupId, _ := utils.GetIntegerInput()

	fmt.Println("Select the userId from below")

	app.UserController.ListAllUsers()

	userId, _ := utils.GetIntegerInput()

	userToBeAdded := app.UserController.GetUser(userId)

	app.groupController.AddUserToGroup(groupId, userToBeAdded)

}

func (app *splitwise) PrintBalanceSheet() {

	app.ExpenseController.PrintBalanceSheet()
}

func (app *splitwise) AddExpense() error {

	fmt.Println("Enter the option\n 0 for group expense \n 1 for non group expense")

	option, ok := utils.GetIntegerInput()

	if ok != nil {
		fmt.Println("Try again by selecting Valid Option")
		return ok
	}

	switch expenseType(option) {
	case groupExpense:
		app.CreateGroupExpense()
	default:
		return fmt.Errorf("Try again by selecting valid option")
	}

	return nil

}

func (app *splitwise) CreateGroupExpense() error {
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

	paidUser := app.UserController.GetUser(paidUserId)

	fmt.Println("Enter the total Amount to be split ")

	totalAmountPaid, err := utils.GetIntegerInput()

	if err != nil {
		return err
	}

	splitAmountsList, err := app.groupController.AddSplitsToUsersFromGroup(groupId)

	if err != nil {
		return err
	}

	expense := app.ExpenseController.AddNewExpense("Expense Name bla bla", paidUser, splitAmountsList, totalAmountPaid, split.UnequalSplit)

	app.groupController.AddExpenseToGroup(groupId, expense)

	return nil
}

func (app *splitwise) CreateNonGroupExpense() error {

	fromUser, toUser, amount, err := app.UserController.GetNewIndividualSplitInput()
	if err != nil {
		return err
	}
	app.ExpenseController.AddIndividualUserExpense("Individual Expense", fromUser, toUser, amount)

	return nil
}
func (app *splitwise) SimplifyDebt() {
	app.ExpenseController.SimplifyDebt()
}
