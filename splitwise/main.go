package main

import (
	"fmt"
	"splitwise/internal/utils"
	"splitwise/splitwise"
)

func main() {

	splitWiseApp := splitwise.NewSplitwiseApp()

	for {

		fmt.Println("Select one of the following options:" +
			"\n 1 to create a user " +
			"\n 2 to create group " +
			"\n 3 to add user to group" +
			"\n 4 to create group expense" +
			"\n 5 to create nongroup expense" +
			"\n 6 to print balance sheet" +
			"\n 7 to simplify debt" +
			"\n 8 to quit")
		option, err := utils.GetIntegerInput()

		if err != nil {
			fmt.Println(err)
			break
		}

		switch option {
		case 1:
			splitWiseApp.AddNewUser()
		case 2:
			splitWiseApp.AddNewGroup()
		case 3:
			splitWiseApp.AddUserToGroup()
		case 4:
			splitWiseApp.CreateGroupExpense()
		case 5:
			splitWiseApp.CreateNonGroupExpense()
		case 6:
			splitWiseApp.PrintBalanceSheet()
		case 7:
			splitWiseApp.SimplifyDebt()
		case 8:
			fmt.Println("Exit Program")
			return
		default:
			fmt.Println("Select Valid Option")
		}
	}
}
