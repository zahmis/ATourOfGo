package main

import (
	"fmt"
	"math/rand"
)


func main() {
	v := rand.Intn(6)
	switch v {
	case 1:
		fmt.Println("凶")
	case 2,3:
		fmt.Println("吉")
	case 4,5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}
