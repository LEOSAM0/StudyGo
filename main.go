package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

/*
Допустим мы с вами описываем конкурс красоты собак
Пускай есть структура собаки, с полями
Строковое имя собаки
Целочисленная судейская оценка
Булевское значение, обозначающее была ли собака вежлива с судьями
Нужно создать несколько собак
Если собака была вежлива с судьями, то добавить собаке один балл судейской оценки
Подобрать наиболее подходящую структуру данных и вывести собак до применения бонусной оценки и после
*/

type Dog struct {
	name       string
	judgeGrade int
	isPolitely bool
}

func (D *Dog) ifPolitely() int {
	if D.isPolitely {
		D.judgeGrade += 1
		return D.judgeGrade
	}
	return D.judgeGrade
}

func main() {
	Funtik := Dog{
		name:       "Funtik",
		judgeGrade: 6,
		isPolitely: true,
	}
	pp.Println(Funtik)

	David := Dog{
		name:       "David",
		judgeGrade: 5,
		isPolitely: true,
	}
	pp.Println(David)

	Hrrr := Dog{
		name:       "Hrrr",
		judgeGrade: 2,
		isPolitely: false,
	}
	pp.Println(Hrrr)

	dogs := map[bool]int{
		Funtik.isPolitely: Funtik.ifPolitely(),
		David.isPolitely:  David.ifPolitely(),
		Hrrr.isPolitely:   Hrrr.ifPolitely(),
	}

	pp.Println(dogs)
	pp.Println(len(dogs))

	i, ok := dogs[Funtik.isPolitely]
	if ok {
		pp.Println(i)
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
