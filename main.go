// GO Slot Machine

package main

import (
	"fmt"
)

// Check if the player wins
func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{} // Will return the multiplier that was achieved on each line E.g. {5, 0, 2} would mean the player won 5x on the first line, 0x on the second and 2x on the third.

	for _, row := range spin { // Loop through each row and check if all symbols are the same
		win := true                      // Set win to true, we then try to disprove this in the next 5 lines
		checkSymbol := row[0]            // Check the first symbol in the row
		for _, symbol := range row[1:] { // Look through subsequent symbols comparing them to checkSymbol (the 1st in the row). [1:] is used as we dont need to compare the first symbol to itself,we start from the 2nd element.
			if checkSymbol != symbol { // As soon as a symbol is different
				win = false // Set win to false
				break       // And break out of the loop
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol]) // If win is still true, append the multiplier for the symbol we won with to the lines slice
		} else {
			lines = append(lines, 0) // If win is false, append 0 to the lines slice as we didn't win on this line.
		}
	}
	return lines
}

// Main function
func main() {
	fmt.Println("Welcome to the Slot Machine!")

	symbols := map[string]uint{ // Key is of type string, Value is of type uint
		"A": 4, // Number denotes the amount of time a symbol will appear in a reel.
		"B": 7,
		"C": 12,
		"D": 20,
	}

	multipliers := map[string]uint{ // Key is of type string, Value is of type uint,
		"A": 20, // Number denotes the multiplier for each symbol.
		"B": 10,
		"C": 5,
		"D": 2,
	}

	var balance uint = 200 // Set the player's balance to 200, Generous right? :D
	row := 3               // Number of rows in the grid
	col := 3               // Number of columns in the grid

	symbolArr := PopulateSymbolArray(symbols) // Generate the grid

	GetName() // Get the player's name

	for balance > 0 { // While the player has money
		bet := GetBet(balance) // Get the player's bet

		if bet == 0 {
			break // If the player enters 0, exit the loop/game
		}

		balance -= bet                          // Deduct the bet from the balance
		spin := CreateGrid(symbolArr, row, col) // Create the grid
		PrintSpin(spin)                         // Print the Slot machine!
		winningLines := checkWin(spin, multipliers)

		for i, multi := range winningLines { // Need to tell user if they won or lost and add the winnings to their balance.
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("You won $%d (%dx) on line #%d\n", win, multi, i+1) // Tell user how much they won and on which line
			}
		}
	}

	fmt.Printf("Thanks for playing! (Balance: $%d)", balance)
}

// Need to run the package itself to see the output.
// Option 1: go run main.go spin.go utils.go (Run all files in one go)
// Option 2: go mod init slot-machine (Create a module)(Run once, will create a go.mod file)
// Can then use: go run . (Run the module) (Everytime you want to run the program)
