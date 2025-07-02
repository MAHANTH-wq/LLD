package balanceSheet_test

import (
	"fmt"
	"splitwise/internal/balanceSheet"
	"splitwise/internal/transaction"
	"splitwise/internal/user"
	"testing"
)

func TestDFSFunction(t *testing.T) {

	balanceSheetController := balanceSheet.NewBalanceSheetController()

	usersList := make([]*user.User, 0)
	balancesList := make([]int, 0)
	transactionList := make([]*transaction.Transaction, 0)
	minTransactionList := make([]*transaction.Transaction, 0)

	userZero := user.NewUser(0, "zero")
	userOne := user.NewUser(1, "one")
	userTwo := user.NewUser(2, "two")
	userThree := user.NewUser(3, "three")
	userFour := user.NewUser(4, "four")

	usersList = append(usersList, userZero, userOne, userTwo, userThree, userFour)
	balancesList = append(balancesList, 0, -500, 400, -100, 200)

	balanceSheetController.ProcessSimplifyDebtUsingDFS(0, balancesList, usersList, transactionList, &minTransactionList)

	for _, value := range minTransactionList {

		fmt.Println("From User, To User, Amount", value.GetFromUser(), value.GetToUser(), value.GetAmount())
	}
	t.Log("Success")

	t.Logf("Length of the list %d", len(minTransactionList))

}

func TestUserListOfBalances(t *testing.T) {

}
