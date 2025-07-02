package splitwise_test

import (
	"fmt"
	"splitwise/splitwise"
	"testing"
)

func TestSplitwiseApp(t *testing.T) {
	app := splitwise.NewSplitwiseApp()

	if app == nil {
		t.Logf("Err creating app")
		return
	}

	app.UserController.AddNewUser("user zero")
	app.UserController.AddNewUser("user one")
	app.UserController.AddNewUser("user two")
	app.UserController.AddNewUser("user three")
	app.UserController.AddNewUser("user four")

	userZero := app.UserController.GetUser(1)
	userOne := app.UserController.GetUser(2)
	userTwo := app.UserController.GetUser(3)
	userThree := app.UserController.GetUser(4)
	userFour := app.UserController.GetUser(5)

	app.ExpenseController.AddNewUser(userZero)
	app.ExpenseController.AddNewUser(userOne)
	app.ExpenseController.AddNewUser(userTwo)
	app.ExpenseController.AddNewUser(userThree)
	app.ExpenseController.AddNewUser(userFour)

	app.ExpenseController.AddIndividualUserExpense("Expense One", userZero, userTwo, 200)
	app.ExpenseController.AddIndividualUserExpense("Expense Two", userFour, userZero, 200)
	app.ExpenseController.AddIndividualUserExpense("Expense Three", userFour, userOne, 600)
	app.ExpenseController.AddIndividualUserExpense("Expense Four", userTwo, userFour, 600)
	app.ExpenseController.AddIndividualUserExpense("Expense Five", userZero, userThree, 100)
	app.ExpenseController.AddIndividualUserExpense("Expense Six", userOne, userZero, 100)

	app.ExpenseController.PrintBalanceSheet()

	fmt.Println("Simplifying the debts")
	app.ExpenseController.SimplifyDebt()

	t.Log("Completed")

}
