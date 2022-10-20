package main

import "fmt"

type Number interface {
	int64 | float64
}

type machine[X Number] struct {
	Data X
}

// type T any - array of any []T
type CustomCollection[T any] []T

type vehicleUpgrader[C int64, T float64] interface {
	Upgrade(C) T
}

func (m machine[X]) Update(C int64) float64 {
	return 1.0
}

func Print[T any](data []T) {
	for _, d := range data {
		fmt.Println(d)
	}
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
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

type mesin[C int64, T float64] struct {
	Data C
}

func (m mesin[C, T]) Upgrade(in C) T {
	return 1.1
}

func main() {

	msn := &mesin[int64, float64]{Data: 2}
	fmt.Println(msn.Upgrade(2))

	m := machine[int64]{Data: 100}
	fmt.Println(m)

	fmt.Println(m.Update(8))

	// Print([]string{"test1", "test2", "test3"})
	// Print([]int{1, 2, 3})

	// https://www.infoworld.com/article/3646036/get-started-with-generics-in-go.html
	// tentukan dulu type nya - [int32]
	// baru isinya based on type - []int32{2, 3, 4}
	x := CustomCollection[int32]([]int32{2, 3, 4})
	fmt.Println(x)

	y := CustomCollection[float64]([]float64{1, 3, 5.2})
	fmt.Println(y)

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
