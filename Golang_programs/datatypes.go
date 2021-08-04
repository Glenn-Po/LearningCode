//datatypes.go

package main

import (
	"fmt"
)

func main() {
	var myage int = 16
	var myname string = "JOHNN GLENN the P.O."
	var is_old bool = false
	n := 23
	lang_point := make(map[string]int) //a map
	lang_point["C"] = 15
	lang_point["C++"] = 16
	lang_point["Java"] = 16
	lang_point["Python"] = 17
	lang_point["Golang"] = 15
	lang_point["Kotlin"] = 14
	lang_point["JavaScript"] = 17
	lang_point["PHP"] = 14
	lang_point["R"] = 14
	lang_point["Swift"] = 14
	lang_point["Ruby"] = 13
	lang_point["Arduino"] = 12
	var arr [5]int                      //fixed size array
	arry := [...]int8{1, 2, 3, 4, 5, 6} //size of array is determined by the compiler at compile-time
	arr[1] = 234
	arr[4] = n
	myarr := make([]int, 3, 100) //make(type, length, capacity), capacity functionis only avalable for slices
	arr2 := []int{12, 56, 9}     //slices are dynamic arrays in effect
	arr2 = append(arr2, 15)      //can add elements to the end of the array
	copy(myarr, arr2)
	fmt.Println("Hello, I am ", myname, ". I am ", myage, is_old)
	fmt.Println("I am HAPPILY learning GOLANG-->really")
	fmt.Printf("%V\n", lang_point)
	fmt.Println(arr, "----", arr2, cap(arr2), "----", arry, "|----|", myarr)
}
