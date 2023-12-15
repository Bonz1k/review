package main

import "fmt"

func main() {
	fmt.Println(comparison(12, 12))
}
func comparison(a, b int) bool {
	if a&b == a {
		return true
	} else {
		return false
	}
}
