package main

import "fmt"

func SumInts(m map[string]int) int {
	var s int
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float32) float32 {
	var s float32
	for _, v := range m {
		s += v
	}

	return s
}

func SumIntsOrFloats[K comparable, V int | float32](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int | float32
}

func main() {
	m1 := map[string]int{
		"first": 34,
		"second": 12,
	}

	m2 := map[string]float32{
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(m1), SumFloats(m2))
	fmt.Printf("Generic Sums: %v and %v infer1: %v and infer2: %v\n", SumIntsOrFloats[string, int](m1), SumIntsOrFloats[string, float32](m2), SumIntsOrFloats(m1), SumIntsOrFloats(m2))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers(m1), SumNumbers(m2))
}
