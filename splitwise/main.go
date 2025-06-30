package main

import (
	"fmt"
	"splitwise/internal/utils"
)

func main() {

	splitWiseApp := newSplitwiseApp()

	for {

		fmt.Println("Select one of the following options:" +
			"\n 1 to create a user " +
			"\n 2 to create group " +
			"\n 3 to add user to group" +
			"\n 4 to create group expense" +
			"\n 5 to create nongroup expense" +
			"\n 6 to print balance sheet" +
			"\n 7 to quit")
		option, err := utils.GetIntegerInput()

		if err != nil {
			fmt.Println(err)
			break
		}

		switch option {
		case 1:
			splitWiseApp.addNewUser()
		case 2:
			splitWiseApp.addNewGroup()
		case 3:
			splitWiseApp.addUserToGroup()
		case 4:
			splitWiseApp.createGroupExpense()
		case 5:
			splitWiseApp.createNonGroupExpense()
		case 6:
			splitWiseApp.printBalanceSheet()
		case 7:
			return
		default:
			fmt.Println("Select Valid Option")
		}
	}
}
