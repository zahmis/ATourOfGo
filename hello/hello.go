package main

import (
	"fmt"
	"math"
	// "math/rand"
	"runtime"
	// "time"
)
import "rsc.io/quote"

func pow(x, n, lim float64) float64 {
	defer fmt.Println("pow")
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {

	defer fmt.Println("world")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}


	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
