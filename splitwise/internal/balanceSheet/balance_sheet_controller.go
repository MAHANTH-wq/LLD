package balanceSheet

import (
	"fmt"
	"splitwise/internal/transaction"
	"splitwise/internal/user"
)

type BalanceSheetController struct {
	userToBalanceSheetMap map[*user.User]*BalanceSheet
}

func NewBalanceSheetController() *BalanceSheetController {
	return &BalanceSheetController{
		userToBalanceSheetMap: make(map[*user.User]*BalanceSheet, 0),
	}
}

func (bs *BalanceSheetController) PrintBalanceSheet() {

	for user, balanceSheet := range bs.userToBalanceSheetMap {

		fmt.Println("Balance Sheet for User with id and name ", user.GetUserId(), user.GetUserName())
		balanceSheet.printBalanceSheet()
	}
}

func (bs *BalanceSheetController) AddNewUser(u *user.User) {
	bs.userToBalanceSheetMap[u] = getNewBalanceSheet()
}

func (bs *BalanceSheetController) AddTransaction(t *transaction.Transaction) {

	bs.userToBalanceSheetMap[t.GetFromUser()].addGetBackMoneyFromUser(t.GetToUser(), t.GetAmount())
	bs.userToBalanceSheetMap[t.GetToUser()].addGiveBackMoneyToUser(t.GetFromUser(), t.GetAmount())
}

func (bs *BalanceSheetController) SimplifyDebt() {

	usersList, balancesList := bs.GetUserListOfBalances()
	fmt.Println(len(usersList))
	fmt.Println(len(balancesList))

	transactionList := make([]*transaction.Transaction, 0)
	minTransactionList := make([]*transaction.Transaction, 0)

	bs.ProcessSimplifyDebtUsingDFS(0, balancesList, usersList, transactionList, &minTransactionList)

	bs.resetUserToBalanceSheetsMap()

	fmt.Println("Print the next transactions")
	for _, transaction := range minTransactionList {
		fmt.Println(transaction.GetFromUser(), transaction.GetToUser(), transaction.GetAmount())
		bs.AddTransaction(transaction)
	}

	fmt.Println("After Simplifying debts , the balance sheet is as follows:")
	bs.PrintBalanceSheet()

}

func (bs *BalanceSheetController) ProcessSimplifyDebtUsingDFS(index int, balancesList []int, usersList []*user.User, transactionsList []*transaction.Transaction, minTransactionList *([]*transaction.Transaction)) {

	if index == len(balancesList) {

		if (len(transactionsList) < len(*minTransactionList)) || (len(*minTransactionList) == 0) {
			*minTransactionList = transactionsList
		}
		return

	}
	if balancesList[index] == 0 {
		bs.ProcessSimplifyDebtUsingDFS(index+1, balancesList, usersList, transactionsList, minTransactionList)
	}

	for i := index + 1; i < len(balancesList); i++ {

		if balancesList[i]*balancesList[index] < 0 {
			balancesList[i] = balancesList[i] + balancesList[index]

			var newTransaction *transaction.Transaction
			if balancesList[index] > 0 {
				newTransaction = transaction.NewTransaction(usersList[index], usersList[i], balancesList[index])

			} else {
				newTransaction = transaction.NewTransaction(usersList[i], usersList[index], -1*balancesList[index])
			}

			transactionsList = append(transactionsList, newTransaction)

			bs.ProcessSimplifyDebtUsingDFS(index+1, balancesList, usersList, transactionsList, minTransactionList)

			balancesList[i] = balancesList[i] - balancesList[index]
			transactionsList = transactionsList[:len(transactionsList)-1]

		}
	}
}

func (bs *BalanceSheetController) GetUserListOfBalances() ([]*user.User, []int) {

	usersList := make([]*user.User, 0)
	balancesList := make([]int, 0)
	for user, balanceSheet := range bs.userToBalanceSheetMap {

		balance := balanceSheet.totalGetBack + balanceSheet.totalGiveBack
		usersList = append(usersList, user)
		balancesList = append(balancesList, balance)
		fmt.Println("User: ", user.GetUserName(), " balance: ", balance)
	}

	return usersList, balancesList
}

func (bs *BalanceSheetController) resetUserToBalanceSheetsMap() {

	for user, _ := range bs.userToBalanceSheetMap {
		bs.userToBalanceSheetMap[user] = getNewBalanceSheet()
	}
}
