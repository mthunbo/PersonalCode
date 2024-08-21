package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func addition(x ...float64) float64 {
	sum := 0.0
	for _, i := range x {
		sum += i
	}
	return sum
}

func subtraction(x ...float64) float64 {
	sum := 0.0
	for _, i := range x {
		sum -= i
	}
	return sum
}

func Sqrt(x float64) float64 {
	z := x / 2
	for z1 := 0.0; z1 != z; {
		z1 = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func operationChooser() {
	choice := 0
	var opscount int
	values := []float64{}
	fmt.Println("Choose what operation to perform by inputing the corresponding number")
	fmt.Print(`
	1. Addition
	2. Subtraction
	3. Multiplication
	4. Division
	5. Square root
	6. Power
	-`)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		clearTerminal()
		fmt.Print(`
		---Addition---
		Please input the amount numbers you want to add together:`)
		fmt.Scan(&opscount)
		fmt.Print("Please input the numbers to add together, press enter after each number\n")
		for i := 0; i < opscount; i++ {
			values = append(values, 0)
			fmt.Scan(&values[i])
		}
		fmt.Printf("\nThe sum of the input is: %f", addition(values...))
	case 2:
		clearTerminal()
		fmt.Print(`
		---Subtraction---
		Please input the amount numbers you want to subtract:`)
		fmt.Scan(&opscount)
		fmt.Print("Please input the numbers to subtract, press enter after each number\n")
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

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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
