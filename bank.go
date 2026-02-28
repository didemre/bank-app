package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit money")
	fmt.Println("3.Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Println("How much do you want to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greaer than 0.")
			return
		}

		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greaer than 0.")
			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		fmt.Println("Goodbye!")

	}

}
