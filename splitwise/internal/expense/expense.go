package expense

import (
	"splitwise/internal/split"
	"splitwise/internal/user"
)

type Expense struct {
	name        string
	paidByUser  *user.User
	splitList   []*split.Split
	totalAmount int
	splitType   split.SplitType
}

func CreateExpense(name string, paidByUser *user.User, splitList []*split.Split, totalAmount int, splitType split.SplitType) *Expense {

	return &Expense{
		name:        name,
		paidByUser:  paidByUser,
		splitList:   splitList,
		totalAmount: totalAmount,
		splitType:   splitType,
	}
}

func (e *Expense) GetSplitList() []*split.Split {

	return e.splitList
}

func (e *Expense) GetPaidByUser() *user.User {
	return e.paidByUser
}

func (e *Expense) GetTotalAmount() int {
	return e.totalAmount
}

func (e *Expense) GetSplitType() split.SplitType {
	return e.splitType
}
