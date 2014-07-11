package main

import "fmt"

func gen() <-chan []int {
	c := make(chan []int)

	go func(c chan []int) {
		defer close(c)
		s := []int{0, 1, 2, 3}

		for i := 0; i < len(s); i++ {
			s[i] = -1
			c <- s
		}
	}(c)

	return c

}

func main() {
	for s := range gen() {
		fmt.Println(s)
	}
}
