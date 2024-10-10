package main

import (
	"fmt"
	"os"
	"strconv"
)

func iter(b int, c chan int) {
	for i:=0; i<b; i++ {
		c <- i
	}
	close(c)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	a := make(chan int)
	b := make(chan int)
	go iter(n, a)
	go iter(n, b)
	for a != nil || b != nil {
		select {
		case x, ok := <- a:
			if ok {
				fmt.Println(x)
			} else {
				a = nil
			}
		case y, ok := <- b:
			if ok {
				fmt.Println(y)
			} else {
				b = nil
			}
		}

	}
}