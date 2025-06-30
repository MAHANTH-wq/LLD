package balanceSheet

import (
	"fmt"
	"splitwise/internal/user"
)

type BalanceSheet struct {

	// Get Back Money is always positive
	totalGetBack int

	//Give Back Money is always negative
	totalGiveBack int

	// If the amount integer is negative it means the user has to giveback that amount to the user in map and
	// if the amount integer is positibe it means the user has to getback that amount to the user in map
	usersMoneyMap map[*user.User]int
}

func getNewBalanceSheet() *BalanceSheet {

	return &BalanceSheet{
		totalGetBack:  0,
		totalGiveBack: 0,
		usersMoneyMap: make(map[*user.User]int, 0),
	}
}

func (bs *BalanceSheet) printBalanceSheet() {
	fmt.Println("Total Money user should get back", bs.totalGetBack)
	fmt.Println("Total Money user should give back", bs.totalGiveBack)

	fmt.Println("Split with individual user")

	for user, money := range bs.usersMoneyMap {
		fmt.Println("Friend User with id and name and total amount", user.GetUserId(), user.GetUserName(), money)
	}

}

func (bs *BalanceSheet) addGetBackMoneyFromUser(u *user.User, amount int) {

	currentAmount, ok := bs.usersMoneyMap[u]

	if !ok {
		bs.usersMoneyMap[u] = amount
		return
	}

	bs.usersMoneyMap[u] = currentAmount + amount
	bs.totalGetBack = bs.totalGetBack + amount
}

func (bs *BalanceSheet) addGiveBackMoneyToUser(u *user.User, amount int) {

	currentAmount, ok := bs.usersMoneyMap[u]

	if !ok {
		bs.usersMoneyMap[u] = -amount
		return
	}
	bs.usersMoneyMap[u] = currentAmount - amount
	bs.totalGiveBack = bs.totalGiveBack - amount
}
