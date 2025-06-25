package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	vm := initVendingMachine()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Click the following buttons on vending machine")
		fmt.Println("1 for requestItem, 2 for insertMoney, 3 for Dispense Items, 4 for remove Items")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println(err)
		}

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Selected Option: ", option)
		switch option {
		case 1:
			vm.requestItem()
		case 2:
			vm.insertMoney()
		case 3:
			vm.dispenseItem()
		case 4:
			vm.removeItem()

		}

	}
}
