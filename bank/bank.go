package main

import "fmt"

func main() {
	var accountBalance float64 = 10000
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	checkBalance := choice == 1

	if checkBalance {
		fmt.Println("Your balance is:", accountBalance )
	}else if choice == 2 {
		fmt.Println("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be more than 0")
			return
		}
		accountBalance += depositAmount
		fmt.Println("balance updated! New amount:", accountBalance)
	}else if choice == 3 {
		fmt.Println("Withdrawl amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("balance updated! New amount:", accountBalance)
	}else{
		fmt.Println("Goodbye!")
	}
}
