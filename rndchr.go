package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Define the --lower flag.
	lower := flag.Bool("lower", false, "Use lowercase (if set)")

	// Default to 8 picks.
	picks := 8

	// Collect arguments for flag parsing, while extracting a numeric argument first.
	newArgs := []string{}
	foundPicks := false

	for _, arg := range os.Args[1:] {
		if !foundPicks {
			if val, err := strconv.Atoi(arg); err == nil {
				picks = val
				foundPicks = true
				continue
			}
		}
		newArgs = append(newArgs, arg)
	}

	// Parse any remaining args for --lower.
	flag.CommandLine.Parse(newArgs)

	// If extra non-flag arguments remain, it's an error.
	if len(flag.Args()) > 0 {
		fmt.Println("Error: Too many arguments.")
		os.Exit(1)
	}

	// Character set: always uppercase, converted to lowercase if needed.
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < picks; i++ {
		ch := string(chars[rand.Intn(len(chars))])
		if *lower {
			ch = strings.ToLower(ch)
		}
		fmt.Print(ch)
	}
	fmt.Println()
}
