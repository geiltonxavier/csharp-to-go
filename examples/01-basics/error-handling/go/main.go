package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrDivideByZero = errors.New("divisão por zero")

func parseOrErr(value string) (int, error) {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("falha ao converter %q: %w", value, err)
	}
	return num, nil
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não foi possível dividir %d por %d: %w", a, b, ErrDivideByZero)
	}
	return float64(a) / float64(b), nil
}

func main() {
	if n, err := parseOrErr("42"); err == nil {
		fmt.Printf("parsed number: %d\n", n)
	} else {
		fmt.Println("erro:", err)
	}

	if n, err := parseOrErr("abc"); err == nil {
		fmt.Printf("parsed number: %d\n", n)
	} else {
		fmt.Println("erro:", err)
	}

	if r, err := divide(10, 2); err == nil {
		fmt.Printf("10 / 2 = %.2f\n", r)
	} else {
		fmt.Println("erro:", err)
	}

	if _, err := divide(10, 0); err != nil {
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("erro de divisão por zero:", err)
		} else {
			fmt.Println("erro:", err)
		}
	}
}
