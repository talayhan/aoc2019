package main

import "fmt"

func is_never_decrease(i int) bool {
	var ret bool

	f := (i % 1000000) / 100000
	e := (i % 100000) / 10000
	d := (i % 10000) / 1000
	c := (i % 1000) / 100
	b := (i % 100) / 10
	a := i % 10

	if a >= b && b >= c && c >= d && d >= e && e >= f {
		ret = true
	} else {
		ret = false
	}

	return ret
}

func check_two_adj_is_same(i int) bool {
	var ret bool

	f := (i % 1000000) / 100000
	e := (i % 100000) / 10000
	d := (i % 10000) / 1000
	c := (i % 1000) / 100
	b := (i % 100) / 10
	a := i % 10

	if f == e || e == d || d == c || c == b || b == a {
		ret = true
	} else {
		ret = false
	}

	return ret
}

func check_extended_adj_rule(i int) bool {
	var ret bool
	//TODO

	return ret
}

func main() {
	total_valid_pas_p1 := 0
	total_valid_pas_p2 := 0
	for n := 124075; n <= 580769; n++ {
		if is_never_decrease(n) && check_two_adj_is_same(n) {
			//fmt.Println("Valid:", n)
			total_valid_pas_p1++
			if check_extended_adj_rule(n) {
				total_valid_pas_p2++
				fmt.Println("Valid part2:", n)
			}
		}
	}

	fmt.Println("Part 1 Answer:", total_valid_pas_p1)
	fmt.Println("Part 2 Answer:", total_valid_pas_p2)
}
