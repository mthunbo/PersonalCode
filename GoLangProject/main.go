package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Addition function that returns the sum of a given number of values
func addition(x ...float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	var sum float64
	for _, i := range x {
		sum += i
	}
	return sum
}

// Subtraction functions that takes an arbitrary number of values adn subtracts from the first given value
func subtraction(x ...float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	sum := x[0]
	for _, i := range x[1:] {
		sum -= i
	}
	return sum
}

// Square root function that returns the square root of a given number
func Sqrt(x float64) float64 {
	sum := x / 2
	for i := 0.0; i != sum; {
		i = sum
		sum -= (sum*sum - x) / (2 * sum)
	}
	return sum
}

// Gets user input and performs the corresponding calculations with it
func getUserInput(choice int) []float64 {
	var opscount int
	values := []float64{}
	operations := [5]string{"ERROR!!!", "Addition", "Subtraction", "Multiplication", "Division"}
	operands := [5]string{"ERROR!!!", "add", "subtract", "multiply", "divide"}

	clearTerminal()
	fmt.Printf("---%s---\nPlease input the amount numbers you want to %s:", operations[choice], operands[choice])
	fmt.Scan(&opscount)
	fmt.Printf("Please input the numbers to %s together, press enter after each number\n", operands[choice])
	//For loop to get the values to calculate on from the user
	for i := 0; i < opscount; i++ {
		values = append(values, 0)
		fmt.Scan(&values[i])
	}
	return values
}

// Function that prints a small terminal menu where the different math operations are available
func operationChooser() {
	choice := 0
	var opscount int
	values := []float64{}
	fmt.Print(`Choose what operation to perform by inputing the corresponding number
	1. Addition
	2. Subtraction
	3. Multiplication
	4. Division
	5. Square root
	6. Power
	-`)
	fmt.Scan(&choice)
	//Switch to choose the operation to perform
	switch choice {
	case 1:
		values = getUserInput(choice)
		fmt.Printf("\nThe sum of the input is: %f", addition(values...))
	case 2:
		clearTerminal()
		fmt.Print(`
		---Subtraction---
		Input the amount numbers you want to subtract:`)
		fmt.Scan(&opscount)
		fmt.Print("Input the numbers to subtract, press enter after each number\n")
		for i := 0; i < opscount; i++ {
			values = append(values, 0)
			fmt.Scan(&values[i])
		}
		fmt.Printf("\nThe sum of the input is: %f", subtraction(values...))

	case 3:
		fmt.Println("")
	case 4:
		fmt.Println("")
	case 5:
		fmt.Println("")
	case 6:
		fmt.Println("")
	default:
		fmt.Println("No choice was made, try again")
		operationChooser()
	}
}

// Function that sends a command(name) to the terminal with possible arguments(args)
func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Switch function that checks to see what OS the user is on to use the corresponding command and arguments for it
func clearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func main() {
	operationChooser()
}
