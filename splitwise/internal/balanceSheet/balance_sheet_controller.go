package balanceSheet

import (
	"fmt"
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

func (bs *BalanceSheetController) AddTransaction(fromUser *user.User, toUser *user.User, amount int) {

	bs.userToBalanceSheetMap[fromUser].addGetBackMoneyFromUser(toUser, amount)
	bs.userToBalanceSheetMap[toUser].addGiveBackMoneyToUser(fromUser, amount)
}
