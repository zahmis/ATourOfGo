package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

}