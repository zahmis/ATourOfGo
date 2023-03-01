package main

import (
	"fmt"
)


func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(CheckOdd(i))
	}
}

func CheckOdd(i int) bool{
	return i%2 == 0
  }
  