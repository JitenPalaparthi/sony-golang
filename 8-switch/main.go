package main

import (
	"fmt"
	"strings"
)

func main() {

	var day uint8 = 4

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("no day")
	}

	// break is automatically invoked after each case. No need to explicitely write break statement

	sal := 10000.500

	switch { //emtpy switch

	case sal <= 10000:
		fmt.Println("Grade D")

	case sal > 10000 && sal <= 15000:
		fmt.Println("Grade C")

	case sal > 15000 && sal <= 25000:
		fmt.Println("Grade B")

	case sal > 25000:
		fmt.Println("Grade A")

	}

	//char :='A' internally char is taken as rune which is nothing but int32
	char := "@"
	char = strings.ToUpper(char)
	switch char {
	case "A", "E", "I", "O", "U":
		fmt.Println("vovel")
	case "B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z":
		fmt.Println("consonent")
	default:
		fmt.Println("special charcter")
	}

}
