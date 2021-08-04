// fact_com_perm
package main

//program to calculate factorial , combinations and permutations

import (
	"errors"
	"fmt"
)

func main() {
	var number int
	var buff string // the same usage as the other programs ::catch \r
	fmt.Println("Enter a number of intergral values to get its factorial:")
	fmt.Scanf("%d", &number)
	fmt.Scan(buff)
	fmt.Println(number, "! = ", fact(number))
	fmt.Println("6C3 = ", nCr(6, 3))
	fmt.Println("6P3 = ", nPr(6, 3))
}

func fact(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("Factorial is undefined for negative numbers")
	} else if n < 2 { //definition for 0 and 1
		return 1, nil
	} else {
		return n * fact(n-1)
	}
}

func fact2(n int) (int64, error) {
	index := 1
	return_value := 1
	if n < 0 {
		return 0, errors.New("Factorial is undefined for negative numbers")
	}
	for index <= n {
		return_value *= index
		index += 1
	}
	return return_value, nil
}

func nCr(n, r int) (int64, error) {
	if n < 0 || r < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	if r > n {
		return 0, errors.New("Check order of values")
	}
	return fact2(n) / (fact2(r) * fact2(n-r)), nil
}

func nPr(n, r int) (int64, error) {
	if n < 0 || r < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	if r > n {
		return 0, errors.New("Check order of values")
	}
	return int64(fact2(n) / fact2(n-r)), nil
}
