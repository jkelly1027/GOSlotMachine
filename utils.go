package main // Must match the package name in the main.go file

import (
	"fmt"
)

// Function to get the player's name.
func GetName() {
	name := ""
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("\nHello %s, let's play the Slot Machine!\n", name)
}

// Ask user for their bet. Must be less than balance
func GetBet(balance uint) uint {
	var bet uint
	for { // Equivalent to while(true)
		fmt.Printf("Enter your bet (Balance: $%d) or 0 to quit: ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("You don't have enough money to make that bet.")
		} else {
			break // To break out of loop if bet is valid
		}
	}
	return bet
}
