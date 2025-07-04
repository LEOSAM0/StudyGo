package main

import (
	dogs "StudyGo/Dogs"
	"fmt"

	"github.com/k0kubun/pp"
)

/*
Допустим мы с вами описываем автомобильную парковку
На парковке есть несколько мест, каждое обозначено каким-то строковым значением (A4, B8, D32 и так далее)
Каджое парковочное место имеет какую-то свою определённую стоимость за час парковки
Нужно создать все парковочные места с их стоимостью, и вывести на экран только парковочные места стоимостью меньше 500 рублей
Кажому месту стоимостью больше 900 рублей нужно сделать скидку 10% (было 1100 рублей стало 990)
Необходимо выбрать наиболее подходящую для решения задачи структуры данных и реализовать описанные пункты
*/

func main() {

	Funtik := dogs.Dog{
		Name:       "Funtik",
		JudgeGrade: 6,
		IsPolitely: true,
	}
	pp.Println(Funtik)

	David := dogs.Dog{
		Name:       "David",
		JudgeGrade: 5,
		IsPolitely: false,
	}
	pp.Println(David)

	Hrrr := dogs.Dog{
		Name:       "Hrrr",
		JudgeGrade: 2,
		IsPolitely: true,
	}
	pp.Println(Hrrr)

	dogs := []dogs.Dog{Funtik, David, Hrrr}

	fmt.Println("До применения бонуса:")
	pp.Println(dogs)

	for i := range dogs {
		dogs[i].IfPolitely()
	}

	fmt.Println("После применения бонуса:")
	pp.Println(dogs)
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
