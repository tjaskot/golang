package main

import (
	"errors"
	"fmt"
	"math"

	"golang.org/x/example/hello/reverse"
)

var bool_1, bool_2, bool_3 bool

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Opt())

	// Hello("watup")
	// Public_call()
	// show_erur_name, show_erur := private_error("no name?")
	// fmt.Println(show_erur_name, show_erur)
	// show_rev("wazzup")
	// fmt.Println(Int(321))
	fmt.Println(ReverseRunes("test"))
	print_vars(1, 2, 3)
	fmt.Println(math.Sqrt(math.Abs(-7)))
	a, b := swap("foo", "bar")
	fmt.Println(a, b)
}

func Hello(name string) string {
	ret_mes := fmt.Sprintf("Hi, %v. Welcome!", name)
	return ret_mes
}

func private_error(erur string) (string, error) {
	if erur == "" {
		return "", errors.New("not a valid name")
	}

	meswage := fmt.Sprintf("Hi, %v. Real name!", erur)
	return meswage, nil
}

func show_rev(rev_str string) {
	fmt.Println(reverse.String(rev_str))
}

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	// fmt.Println(s)
	r := []rune(s)
	fmt.Println(r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func print_vars(bool_1, bool_2, bool_3 int) {
	fmt.Println(bool_1, bool_2, bool_3)
}

func swap(x, y string) (string, string) { // func swap takes two strings and returns two strings in reverse
	return y, x
}

func add_count() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
