// bubble_sort
package main

import (
	"fmt"
	//"regexp/syntax"
)

func main() {
	array := []int{1, 45, 0, 78, -90, 8, 66, 56, 9}

	fmt.Println("Program to sort a sequence by Insertion sort")
	for i := 0; i < len(array); i++ {
		current := array[i]
		j := i
		for j > 0 && array[j-1] > current {
			array[j] = array[j-1]
			j--
		}
		array[j] = current
	}
	fmt.Println(array)
}
