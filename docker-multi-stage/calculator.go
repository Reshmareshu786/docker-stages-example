package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hi, what is your name? ")
	username, _ := reader.ReadString('\n')

	// Trim the newline character from the username
	username = strings.TrimSpace(username)

	// Greet the user with their name
	fmt.Printf("Hi %s, I am a calculator app ....\n", username)

	for {
		// Read input from the user
		fmt.Print("Enter any calculation (Example: 1 + 2 (or) 2 * 5 -> Please maintain spaces as shown in example): ")
		text, _ := reader.ReadString('\n')

		// Trim the newline character from the input
		text = strings.TrimSpace(text)

		// Check if the user entered "exit" to quit the program
		if text == "exit" {
			fmt.Printf("Goodbye, %s!\n", username)
			break
		}

		// Split the input into parts: the left operand, the operator, and the right operand
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Convert the operands to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Perform the calculation based on the operator
		var result int
		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Cannot divide by zero. Try again.")
				continue
			}
			result = left / right
		default:
			fmt.Println("Invalid operator. Try again.")
			continue
		}

		// Print the result
		fmt.Printf("%s, the result is: %d\n", username, result)
	}
}
