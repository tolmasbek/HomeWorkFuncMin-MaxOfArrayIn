package main

import (
	"fmt"
)

func main() {
	array := [] int64 {16,4,7,9,11,23}

	MinOfArray(array)

	MaxOfArray(array)
}

func MinOfArray(arr []int64) {
	min := arr [0]
	for i := 0; i < len(arr); i++ {
		if min > arr [i] {
			min = arr [i]
		}
	}
	fmt.Println("Минимальный элемент массива = ", min)
}

func MaxOfArray(arr []int64)  {
	max := arr [0]
	for i := 0; i < len(arr); i++ {
		if max < arr [i] {
			max = arr [i]
		}
	}
	fmt.Println("Максимальный элемент массива = ", max)
}