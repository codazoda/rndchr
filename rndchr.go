package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var err error
	// Set of characters
	availableChars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	totalChars := len(availableChars)
	// Check if there is a single command line argument
	argLength := len(os.Args[1:])
	if argLength > 1 {
		fmt.Println("Error: Too many command line arguments.")
		os.Exit(1)
	}
	// Set the defaul number of picks to 8
	numOfPicks := 8
	// If there is one argument, use it
	if argLength == 1 {
		numOfPicks, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: Argument 1 was not a valid number number.")
		}
	}
	// Loop as many times as we like
	for i := 0; i < numOfPicks; i++ {
		// Seed the random number generator
		rand.Seed(time.Now().UnixNano())
		// Generate random digits
		pick := rand.Intn(totalChars)
		// Output this digit
		fmt.Print(string(availableChars[pick]))
	}
	fmt.Println()
}
