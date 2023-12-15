package main

import "fmt"

var num int = 1

func main() {
	fmt.Println(num)
	increaseInt(&num)
	fmt.Println(num)
	increaseInt(&num)
	fmt.Println(num)
	increaseInt(&num)
	fmt.Println(num)
}
func increaseInt(numPointer *int) {
	*numPointer += 1
}
