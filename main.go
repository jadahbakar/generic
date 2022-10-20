package main

import "fmt"

// [T any] : type T any - []T : array of any
type CustomCollection[T any] []T

func main() {
	x := CustomCollection[int32]([]int32{2, 3, 4})
	fmt.Println(x)

	y := CustomCollection[float64]([]float64{1, 3, 5.2})
	fmt.Println(y)
}
