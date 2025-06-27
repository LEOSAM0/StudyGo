package main

import "fmt"

//main
//number 1
//number 2
//oops
//FuckYouaa
//HoBA
//number 1
//number 2
//oops
//stop

func main() {
	fmt.Println("MAIN")

	hello()

	println("EgegeY")

	hi()

	hello()

	println("STOP!")
}

func hello() {
	defer func() {
		fmt.Println("Hello number 2")

		defer func() {
			fmt.Println("OOOPS!")
		}()

	}()

	fmt.Println("Hello number 1")
}

func hi() {
	println("FuckYouaaaiiie")

	defer func() {
		println("HoBA")
	}()
}
