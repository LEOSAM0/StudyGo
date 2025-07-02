package main

import (
	automobile "StudyGo/Automobile"
	greeting "StudyGo/Greeting"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	name string
	age  int
}

func main() {
	greeting.Hi()

	Anton := User{"Anton", 32}
	pp.Println(Anton)

	Audi := automobile.NewAutomobile("Audi", true, 3.21, 234.7, 8)
	pp.Println(Audi)
}

func changePtr(in *int, str *string, bl *bool, fl *float64) {
	if in != nil && str != nil && bl != nil && fl != nil {
		*in = -3
		*str = "world"
		*bl = false
		*fl = -23.1
		fmt.Println("указатели не нулевые")
	} else {
		fmt.Println("NIL!")
	}

}

// defer method (stack)
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

/*
часы 12ти часовой формат
d - часовая и минутная стрелки
0 < d < 360
h - часы
m - минуты
It is h hours m minutes.
*/
func dial() {
	var d uint16
	fmt.Scan(&d)
	h := d / 30
	m := (d % 30) * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}
