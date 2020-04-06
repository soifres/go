package main

import (
	"fmt"
)

func main() {

	// строковая переменная
	var name_one string
	name_one = "\tvariable name_one"
	var name_two = "\tvariable name_two"
	name_four := "\tvariable three"

	var i int = 1
	var q float64 = 15.0
	var is bool = true

	fmt.Println(name_one)
	fmt.Println(name_two)
	fmt.Println(name_four)

	fmt.Println(i)
	fmt.Println(q)
	fmt.Println(is)
	fmt.Println(name_four)

	var result float64 = q / float64(i+i)
	fmt.Println("\tРезультат q / float64(i+i):", result)

	var arr_1 [10]int = [10]int{1, 2, 3, 5, 78, 56}
	for i := 0; i < 10; i++ {
		fmt.Println(i, arr_1[i])
	}

	if do_test("Print") {
		fmt.Println("Test complete")
		fmt.Printf("Second string")
	}

	// http.Listen
}

// do_test Function hhich do something
func do_test(s string) bool {
	return true
}
