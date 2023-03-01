package main

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting/v2"
)

// func SwapPointer(x, y *int) {
//      *x, *y = *y, *x
// }

// func main() {
//      x, y := 10, 20
//      SwapPointer(&x, &y)
//      fmt.Println(x, y)
// }

type MyInt int

func (i *MyInt) Inc() { *i++ }

func main() {
        var n MyInt = 10
        fmt.Println(n)
        n.Inc()
        fmt.Println(n)
        n.Inc()
        fmt.Println(n)

        // fmt.Println(greeting.Do())
        fmt.Println(greeting.Do(time.Now()))
}