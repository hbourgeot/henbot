package calc

import (
	"fmt"
	"strconv"
)

func convert(sum1, sum2 string) (int, int, error) {
	a, err := strconv.Atoi(sum1)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.Atoi(sum2)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

func Add(sumOperation []string) (int, error) {
	a, b, err := convert(sumOperation[1], sumOperation[2])
	if err != nil {
		return 0, err
	}

	return a + b, nil
}

func Sub(subOperation []string) (int, error) {
	a, b, err := convert(subOperation[1], subOperation[2])
	if err != nil {
		return 0, err
	}

	return a - b, nil
}

func Mult(mulOperation []string) (int, error) {
	a, b, err := convert(mulOperation[1], mulOperation[2])
	if err != nil {
		return 0, err
	}

	return a * b, nil
}

func Div(divOperation []string) (int, error) {
	a, b, err := convert(divOperation[1], divOperation[2])
	if err != nil {
		return 0, err
	}

	return a / b, nil
}

func Mod(modOperation []string) (int, error) {
	a, b, err := convert(modOperation[1], modOperation[2])
	if err != nil {
		return 0, err
	}

	return a % b, nil
}

func Perc(percOperation []string) (float64, error) {
	a, err := strconv.ParseFloat(percOperation[1], 64)
	_, b, err := convert(percOperation[1], percOperation[2])
	if err != nil {
		return 0, err
	}
	fmt.Print(a)
	percentage := float64(a / 100)

	return percentage * float64(b), nil
}
