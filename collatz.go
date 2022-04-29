package main

import (
	"flag"
	"fmt"
)

var something int = 21

func main() {
	nFlag := flag.Int("n", something, "text")
	flag.Parse()
	n := *nFlag
	fmt.Println(n)
	for n != 1 {
		if n%2 == 1 {
			n = n*3 + 1
		} else {
			n = n / 2
		}
		fmt.Println(n)
	}
}
