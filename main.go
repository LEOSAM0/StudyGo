package main

import (
	automobile "StudyGo/Automobile"
	greeting "StudyGo/Greeting"
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

func (D *Dog) ifPolitely() {
	if D.isPolitely {
		D.judgeGrade += 1
	}
}

func main() {
	greeting.Hi()

	Audi := automobile.NewAutomobile("Audi", true, 3.21, 234.7, 8)
	pp.Println(Audi)

	mapa1 := map[string]int{
		"A": 54,
		"B": 3,
		"C": 81,
	}
	pp.Println(mapa1)

	mapa1["D"] = 21
	v, ok := mapa1["A"]
	if ok {
		pp.Println(v)
	} else {
		pp.Println("Not found")
	}

	delete(mapa1, "B")
	pp.Println(mapa1)

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
