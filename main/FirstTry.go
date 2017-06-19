package main

import (
	"fmt"
	"strings"
)

func main(){
	// the must have Hello World
	println("Hello World!")
	fmt.Printf("Howdy World!\n")

	stringFunction()

	integerFunction()

}

func stringFunction(){
	var myName = "Florent Haxha";
	// Boolean
	println(strings.Contains(myName, "Florent"))
	// Boolean
	println(strings.ContainsAny(myName, "F"))
	// returns 0 if matches and 1 if it doesn't
	println(strings.Compare(myName, "Florent Haxha"))
	// return index of given symbol
	println(strings.Index(myName, "t"))
	// runs left to right and trims away the substring from the string, only happens if the substring also is the start of the string
	println(strings.Trim(myName, "Flo"))
	// simply prints upper case
	println(strings.ToUpper(myName))
}

func integerFunction(){
	for i := 0; i <= 10 ; i++ {
		if i % 2 == 0 {
			println(i);
		}
	}


}