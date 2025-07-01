package main

import "fmt"

/*
Мы описываем квартиру
У каждой квартиры есть адрес, площадь, этаж, и стоимость
Каждая квартира может менять свою стоимость
Адрес не должен быть пустой, и не может изменяться
Площадь должна быть больше 10 квадратных метров и не может изменяться
Этаж должен быть больше нуля и меньше либо равен ста
Стоимость должна быть больше нуля и может изменяться
Необходимо описать соответсвующую структуру с конструктором, полями и методами
То, что поля на текущем этапе изучение мы сможем изменить напрямую и сделать их невалидными мы во внимание не берём
*/

type Flat struct {
	adress string
	square float64
	floor  int
	price  float64
}

func NewFlat(adr string, sq float64, fl int, pr float64) Flat {
	if adr == "" {
		return Flat{}
	}

	if sq < 10 {
		return Flat{}
	}

	if fl < 0 || fl > 100 {
		return Flat{}
	}

	if pr <= 0 {
		return Flat{}
	}

	return Flat{
		adress: adr,
		square: sq,
		floor:  fl,
		price:  pr,
	}
}

func (F Flat) GetAdress() string {
	return F.adress
}

func (F Flat) GetSquare() float64 {
	return F.square
}

func (F Flat) GetFloor() int {
	return F.floor
}

func (F *Flat) ChangePrice(pr float64) {
	F.price = pr
}

func main() {
	BigHouse := NewFlat("Sulimova", 32.3, 3, 471.8)
	fmt.Println(BigHouse)
	BigHouse.ChangePrice(9999.91)
	fmt.Println(BigHouse)
	floor := BigHouse.GetFloor()
	fmt.Println(floor)
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
