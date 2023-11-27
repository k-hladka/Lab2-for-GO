package main

import (
	"fmt"

	"funcs/funcs"
)

func main() {
	min := funcs.MinValues(200.51, 2.12, 10)
	fmt.Println("Мінімальне значення з трьох чисел = ", min)

	avg := funcs.AvgValues(1, 3, -10.2)
	fmt.Println("Середнє значення трьох чисел = ", avg)

	diff := funcs.DiffEq(1, 3)
	fmt.Println("Диференціальне рівняння першого порядку. Однорідне рівняння вигляду (x+y)/x = ", diff)
}
