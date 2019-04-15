package main

import "fmt"

func pointTest(n *int) {
	*n = 9
}

func main() {
	var count = 3
	var p *int = &count
	fmt.Println(p)
	pointTest(&count)
	fmt.Println(count)
	var p1 *int
	if p1 == nil {
		fmt.Println("count is nil")
	}
}
