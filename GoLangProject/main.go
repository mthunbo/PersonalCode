package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Addition function that returns the sum of a given number of values
func addition(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	var sum float64
	for _, i := range x {
		sum += i
	}
	return sum
}

// Function that takes n numbers and subtracts from the first given value
func subtraction(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	result := x[0]
	for _, i := range x[1:] {
		result -= i
	}
	return result
}

// Function that takes n numbers and multiplies them together
func multiplication(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	var result float64
	for _, i := range x {
		result *= i
	}
	return result
}

// Function that divdes the first given number by the second given
func division(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	result := x[0]
	for _, i := range x[1:] {
		result /= i
	}
	return result
}

// Function that returns the square root of a given number
func root(base, n float64) float64 {
	result := 1.0
	for i := 0; i < 1000; i++ {
		result = ((n-1)*result + base/power(result, n-1)) / n
	}
	return result
}

// Function that returns the n-th power of a given number
func power(base, n float64) float64 {
	result := 1.0
	for i := 0; i < int(n); i++ {
		result *= base
	}
	return result
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
	fmt.Printf("Please input the numbers to %s, press enter after each number\n", operands[choice])
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
	var values []float64
	var base float64
	var nth float64

	fmt.Print(`Choose what operation to perform by inputing the corresponding number
	1. Addition
	2. Subtraction
	3. Multiplication
	4. Division
	5. Square root
	6. Power
	7. End program
	-`)
	fmt.Scan(&choice)
	//Switch to choose the operation to perform
	switch choice {
	case 1:
		values = getUserInput(choice)
		fmt.Printf("\nThe sum of the input is: %f", addition(values))
	case 2:
		values = getUserInput(choice)
		fmt.Printf("\nThe result of the input is: %f", subtraction(values))
	case 3:
		values = getUserInput(choice)
		fmt.Printf("\nThe result of the input is: %f", multiplication(values))
	case 4:
		values = getUserInput(choice)
		fmt.Printf("\nThe result of the input is: %f", division(values))
	case 5:
		fmt.Printf("---Root---\nPlease input the number you wish to get the nth root of, followed by the root coefficient:\n")
		fmt.Scan(&base)
		fmt.Scan(&nth)
		fmt.Printf("\nThe result of the input is: %f", root(base, nth))
	case 6:
		fmt.Printf("---Power---\nPlease input the number you wish to get the nth power of and the number to uplift it with:\n")
		fmt.Scan(&base)
		fmt.Scan(&nth)
		fmt.Printf("\nThe result of the input is: %f", power(base, nth))
	case 7:
		clearTerminal()
		fmt.Println("You sure you want to exit the program? [Yes/No]")

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
