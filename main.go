package main

import (
	"fmt"
	"math/rand"
)

func toString(chararr []string) []string {
	chararray := make([]string, len(chararr))
	for i, val := range chararr {
		// chararray[i] = ((val == '\x00') ? '-' : val)
		if val == " " {
			chararray[i] = "-"
		} else {
			chararray[i] = val
		}
	}

	return chararray
}

func main() {
	// fmt.Print("Hello")

	arr := [...]string{"animal", "happiness", "indefinite", "steady", "birthday", "extreme", "right", "properties", "ceremony", "beneath", "infrmation", "students", "employee"}
	var selectedword string
	var lenofarr int = len(arr) - 1
	randomnum := rand.Intn(lenofarr-0) + 0
	selectedword = arr[randomnum]

	// string to array of characters

	chararray := make([]string, len(selectedword))

	for i, r := range selectedword {
		chararray[i] = string(r)

	}
	cha := toString(chararray)
	fmt.Print(cha)

	// chararray created

}
