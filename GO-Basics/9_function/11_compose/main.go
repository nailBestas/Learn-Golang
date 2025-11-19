package main

import (
	"fmt"
)

// tek sorumlulu funclar
func kare(x int) int       { return x * x }
func ikiyleCarp(x int) int { return x * 2 }

// funclari birlestirelim
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	kareVeIkiylecarp := compose(ikiyleCarp, kare)
	fmt.Println(kareVeIkiylecarp(10))
}
