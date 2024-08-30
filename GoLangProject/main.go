package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	result := 1.0
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
func root(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	result := 1.0
	for i := 0; i < 1000; i++ {
		result = (x[1]-1)*result + x[0]/x[1]
	}
	return result
}

// Function that returns the n-th power of a given number
func power(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	result := 1.0
	for i := 0; i < int(x[1]); i++ {
		result *= x[0]
	}
	return result
}

// Gets user input and performs the corresponding calculations with it
func getUserInput(choice int) []float64 {
	var opscount int
	values := []float64{}
	operations := [7]string{"ERROR!!!", "Addition", "Subtraction", "Multiplication", "Division", "Root", "Power"}
	operands := [7]string{"ERROR!!!", "add", "subtract", "multiply", "divide", "Root", "Power"}

	clearTerminal()
	// If that checks what operations to perform as root and power works a little different than the others
	if choice < 4 && choice > 0 {
		fmt.Printf("---%s---\nPlease input the amount numbers you want to %s:", operations[choice], operands[choice])
		fmt.Scan(&opscount)
		fmt.Printf("Please input the numbers to %s, press enter after each number\n", operands[choice])
		//For loop to get the values to calculate on from the user
		for i := 0; i < opscount; i++ {
			values = append(values, 0)
			fmt.Scan(&values[i])
		}
	} else {
		fmt.Printf("---%s---\nPlease input the base number and the coefficient to %s the base with:\n", operations[choice], operands[choice])
		values = append(values, 0)
		fmt.Scan(&values[0])
		values = append(values, 0)
		fmt.Scan(&values[1])
	}
	return values
}

// Function that checks if user wants to continue er exit program
func askForExit() bool {
	var exit string
	fmt.Scan(&exit)
	exit = strings.ToLower(exit)
	if exit == "yes" || exit == "y" {
		return true
	} else if exit == "no" || exit == "n" {
		return false
	}
	return false
}

// Function that contains the loop that allows for continuous operations
func operationLoop(fn func(x []float64) float64, choice int) {
	var values []float64
	result := 0.0

	for {
		values = getUserInput(choice)
		result = fn(values)
		fmt.Printf("\nThe result of the input is: %f\n", result)
		fmt.Println("Do you want to continue current math operations[Yes] or exit to do others?[No]")
		if !askForExit() {
			operationChooser()
		}
	}
}

// Function that prints a small terminal menu where the different math operations are available
func operationChooser() {
	choice := 0

	clearTerminal()
	fmt.Print(`Choose what operation to perform by inputing the corresponding number
	1. Addition
	2. Subtraction
	3. Multiplication
	4. Division
	5. Square root
	6. Power
	7. Exit program
	-`)
	fmt.Scan(&choice)
	//Switch to choose the operation to perform
	switch choice {
	case 1:
		operationLoop(addition, choice)
	case 2:
		operationLoop(subtraction, choice)
	case 3:
		operationLoop(multiplication, choice)
	case 4:
		operationLoop(division, choice)
	case 5:
		operationLoop(root, choice)
	case 6:
		operationLoop(power, choice)
	case 7:
		clearTerminal()
		fmt.Println("Are you sure you want to exit the program? [Yes/No]")
		if askForExit() {
			os.Exit(0)
		} else {
			operationChooser()
		}
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
