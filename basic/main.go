package main

import "fmt"

func main() {
	total := 0
	for range "abc" {
		total++
	}
	fmt.Print(total)
}
