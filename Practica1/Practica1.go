package main

import (
	"fmt"
)

const PROMPT = ">> "

type operation struct {
	opType string

}

func main(){
	//fmt.Println("Hola")
	//file, err := os.OpenFile("CalcStorage.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer  file.Close()
	//Check(err)
	//file.WriteString("Esto es una prueba")
	//file.WriteString("Esto ya no es una prueba")
	//file.Close()
	var mainOp int = 0
	for mainOp != 3{
		mainOp = MainMenu()
		switch mainOp {
		case 1:
			fmt.Println("Loading calculator menu...")
				var calcOp int = 0
				for calcOp != 5{
					calcOp = CalMenu()
					switch calcOp {
					case 1:
						Sum()
					case 2:
						Subtraction()
					case 3:
						Multiplication()
					case 4:
						Division()
					case 5:
						fmt.Println("Exiting the calculator...")

					}
				}
		case 2:
			fmt.Println("Loading files menu...")
			var fileOp int = 0
			for fileOp != 6{
				fileOp = CalMenu()
				switch fileOp {
				case 1:
					Add()
				case 2:
					Subtraction()
				case 3:
					Multiplication()
				case 4:
					Division()
				case 5:
					fmt.Println("Exiting the calculator...")

				}
			}
		case 3:
			fmt.Println("Exiting the program...")

		}
	}

}

// Utils -------------------------------------------------------------------------------------------------------------->
// IsBetween validates if a number is between an interval
func IsBetween(x, a, b int) bool{

	if a<= x && x <= b{
		return true
	}
	return  false
}

// Check checks for errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Calculator --------------------------------------------------------------------------------------------------------->
func Subtraction() int{
	var x, y int
	fmt.Println("Enter the first number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Println("Enter the second number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&y)
	fmt.Println()
	result := x - y
	answer := fmt.Sprintf("(%d - %d) = %d", x,y, result)
	fmt.Println(answer)
	fmt.Println()
	return x - y
}

func Sum() int{
	var x, y int
	fmt.Println("Enter the first number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Println("Enter the second number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&y)
	fmt.Println()
	result := x + y
	answer := fmt.Sprintf("(%d + %d) = %d", x,y, result)
	fmt.Println(answer)
	fmt.Println()
	return result
}

func Multiplication() int{
	var x, y int
	fmt.Println("Enter the first number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Println("Enter the second number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&y)
	fmt.Println()
	result := x * y
	answer := fmt.Sprintf("(%d * %d) = %d", x,y, result)
	fmt.Println(answer)
	fmt.Println()
	return result
}

func Division() int{
	var x, y int
	fmt.Println("Enter the first number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Println("Enter the second number: ")
	fmt.Print(PROMPT)
	fmt.Scanln(&y)
	fmt.Println()
	result := x / y
	answer := fmt.Sprintf("(%d / %d) = %d", x,y, result)
	fmt.Println(answer)
	fmt.Println()
	return result
}

// Menu --------------------------------------------------------------------------------------------------------------->
func MainMenu() int {
	var op int = 0
	for op == 0{
		fmt.Println("1.Calculator")
		fmt.Println("2.Files")
		fmt.Println("3.Exit")
		fmt.Print(PROMPT)
		fmt.Scanln(&op)
		fmt.Println()
		if !IsBetween(op, 1, 3){
			fmt.Println("Invalid option please try again!")
		}
	}
	return op
}
func CalMenu() int {
	var op int = 0
	for op == 0{
		fmt.Println("1.Sum")
		fmt.Println("2.Subtraction")
		fmt.Println("3.Multiplication")
		fmt.Println("4.Division")
		fmt.Println("5.Exit")
		fmt.Print(PROMPT)
		fmt.Scanln(&op)
		fmt.Println()
		if !IsBetween(op, 1, 5){
			fmt.Println("Invalid option please try again!")
		}
	}
	return op
}

func FileMenu() int {
	var op int = 0
	for op == 0{
		fmt.Println("1.Add")
		fmt.Println("2.Show")
		fmt.Println("3.Search")
		fmt.Println("4.Update")
		fmt.Println("5.Delete")
		fmt.Println("6.Exit")
		fmt.Print(PROMPT)
		fmt.Scanln(&op)
		fmt.Println()
		if !IsBetween(op, 1, 6){
			fmt.Println("Invalid option please try again!")
		}
	}
	return op
}

// Files

func Add(){

}

// Decorators

func Timer(){

}