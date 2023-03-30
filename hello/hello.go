package main

import "fmt"

var n int
var w []string

func init() {
	fmt.Println("init")
	fmt.Scanln(&n)
	w = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
	}
	fmt.Println(w)
}

func main() {
	is := false
	for i := 0; i < n; i++ {
		fmt.Println(w[i])
		if w[i] == "and" || w[i] == "the" || w[i] == "not" || w[i] == "you" || w[i]=="that" {
			is = true
		}
	}

	if is {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	
}
