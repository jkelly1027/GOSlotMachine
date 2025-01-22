package main

import (
	"fmt"
	"math/rand"
)

// Populate the symbol array
func PopulateSymbolArray(symbols map[string]uint) []string { // Key: string, Value: uint, Return: []string (Slice of strings)
	symbolArray := []string{} // Create an empty slice of strings

	for symbol, count := range symbols { // Loop through symbols map, populate slice with symbols
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, symbol) // Append symbol to slice(symbolArray)
		}
	}
	return symbolArray
}

// Function to get a random number between min and max
func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min // Generate a random number between min and max
	return randomNumber
}

// Create the grid
func CreateGrid(reel []string, rows int, cols int) [][]string { // Take in symbols (reel []string), rows and columns, Return a 2d slice of strings
	result := [][]string{} // Create an empty 2d slice of strings, each internal slice represents a row. {{A,B,C},{A,B,C},{A,B,C}}

	// Creating empty grid
	for i := 0; i < rows; i++ {
		result = append(result, []string{}) // Append X number of empty slices/rows of strings to the result slice, adding rows
	}

	// Populating the grid: [[A,,],[A,,],[A,,]] -> [[A,B,],[A,B,],[A,B,]] -> [[A,B,C],[A,B,C],[A,B,C]] Insert at pos 0 of first row, pos 0 of second row, pos 0 of third row, pos 1 of first row...
	for col := 0; col < cols; col++ { //[[COL1],[COL2],[COL3]]
		selected := map[int]bool{}        // Create a map to keep track of the positions that have already been selected from the reel(Symbols) int referring to the index of the reel and bool for true/false where true means selected.
		for row := 0; row < rows; row++ { // [[[C1R1],[C1R2],[C1R3]],[[C2R1],[C2R2],[C2R3]],[[C3R1],[C3R2],[C3R3]]]
			for {
				randomIndex := getRandomNumber(0, len(reel)-1) // Get a random number between 0 and the length of the reel, randomly pick a position from the reel.
				_, exists := selected[randomIndex]             // Check if we have already selected the random index. _ will give us whatever the value is inside "selected[randomIndex]" and exists will tell us if it exists within selected or not. Don't care if its t
				if !exists {
					selected[randomIndex] = true                         // If the random number is not already selected, record it in selected and set it to true
					result[row] = append(result[row], reel[randomIndex]) // Insert the symbol from the reel at the random index into the grid at the current row and column
					break
				}
			}
		}
	}
	return result
}

// Print Spin function
func PrintSpin(spin [][]string) {
	for _, row := range spin { // Loop through all slices in the spin slice.
		for j, symbol := range row { // Then loop through all of the rows in each slice
			fmt.Print(symbol)    // Print the symbol
			if j != len(row)-1 { // But only if not in the last column will we print the separator (A|B|C) Don't want one after
				fmt.Print(" | ")
			}
		}
		fmt.Println() // Print a new line after each row
	}
}
