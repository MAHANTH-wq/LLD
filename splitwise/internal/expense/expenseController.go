package expense

import (
	"splitwise/internal/balanceSheet"
	"splitwise/internal/split"
	"splitwise/internal/transaction"
	"splitwise/internal/user"
)

type ExpenseController struct {
	balanceSheetController *balanceSheet.BalanceSheetController
}

func NewExpenseController() *ExpenseController {

	return &ExpenseController{
		balanceSheetController: balanceSheet.NewBalanceSheetController(),
	}

}

func (ec *ExpenseController) AddNewUser(u *user.User) {
	ec.balanceSheetController.AddNewUser(u)
}

func (ec *ExpenseController) PrintBalanceSheet() {
	ec.balanceSheetController.PrintBalanceSheet()
}

func (ec *ExpenseController) AddNewExpense(name string, paidByUser *user.User, splitList []*split.Split, totalAmount int, splitType split.SplitType) *Expense {

	expense := CreateExpense(name, paidByUser, splitList, totalAmount, splitType)

	ec.ProcessExpense(expense)

	return expense

}

func (ec *ExpenseController) AddIndividualUserExpense(name string, paidByUser *user.User, toUser *user.User, totalAmount int) *Expense {

	splitValue := split.CreateNewSplit(toUser, totalAmount)
	listOfSplits := make([]*split.Split, 0)
	listOfSplits = append(listOfSplits, splitValue)
	expense := CreateExpense(name, paidByUser, listOfSplits, totalAmount, split.UnequalSplit)
	ec.ProcessExpense(expense)
	return expense

}
func (ec *ExpenseController) ProcessExpense(e *Expense) error {

	splitList := e.GetSplitList()
	for _, split := range splitList {

		if split.GetSplitUserId() == e.GetPaidByUser().GetUserId() {
			continue
		}

		newTransaction := transaction.NewTransaction(e.GetPaidByUser(), split.GetSplitUser(), split.GetSplitAmount())

		ec.balanceSheetController.AddTransaction(newTransaction)
	}

	return nil
}

func (ec *ExpenseController) SimplifyDebt() {
	ec.balanceSheetController.SimplifyDebt()
}
