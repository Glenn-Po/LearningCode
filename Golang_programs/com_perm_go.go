// com_perm_go
package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(fact(5))
	fmt.Println(fact(1))
	fmt.Println(fact(-3))

}
func fact(n int64) int64 {
	if n < 0 {
		return nil
	} else if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
