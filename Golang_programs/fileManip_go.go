// fileManip_go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number of rows:")
	var rows int
	fmt.Scanf("%d", &rows)
	for i := 0; i < rows; i++ {
		for space := 0; space < rows-i; space++ {
			fmt.Print("  ")
		}

		var coef int

		for j := 0; j <= i; j++ {
			if j == 0 || i == 0 {
				coef = 1
			} else {
				coef = coef * (i - j + 1) / j
			}
			fmt.Printf("%4d", coef)
		}
		fmt.Println()

	}
}
