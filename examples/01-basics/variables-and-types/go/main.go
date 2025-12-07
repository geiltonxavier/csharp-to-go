package main

import (
	"fmt"
	"strings"
)

func main() {
	var count int = 3
	name := "Go"
	const pi = 3.14
	isReady := true

	numbers := [3]int{1, 2, 3}
	primes := []int{2, 3, 5}
	ages := map[string]int{
		"ana":  29,
		"joao": 31,
	}

	fmt.Printf("count: %d, name: %s, pi: %.2f, ready: %t\n", count, name, pi, isReady)
	fmt.Printf("numbers length: %d, first: %d\n", len(numbers), numbers[0])

	primes = append(primes, 7)
	fmt.Printf("primes: %s\n", strings.Trim(fmt.Sprint(primes), "[]"))

	fmt.Printf("idade de ana: %d\n", ages["ana"])

	var zeroInt int        // zero value is 0
	var emptyString string // zero value is "" not nil
	fmt.Printf("zero int: %d, zero string: %q\n", zeroInt, emptyString)
}
