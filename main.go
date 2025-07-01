package main

import "fmt"

/*
дано трёхзначное число
вывести YES - все цифры отличаются друг от друга
иначе NO
*/

func main() {
	var str string
	fmt.Scan(&str)
	if str[0] == str[1] || str[0] == str[2] || str[1] == str[2] {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
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

//defer method (stack)
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

/*часы 12ти часовой формат
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
