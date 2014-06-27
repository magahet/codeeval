package main

import (
	"fmt"
	"math"
)

func generatePrimes(N int) <-chan int {
	out := make(chan int)
	go func() {
		notPrime := make([]bool, N+1)
		limit := int(math.Sqrt(float64(N)))
		for i := 2; i <= limit; i++ {
			if notPrime[i] == false {
				out <- i
				for j := i * i; j <= N; j += i {
					notPrime[j] = true
				}
			}
		}

		for i := limit + 1; i <= N; i++ {
			if notPrime[i] == false {
				out <- i
			}
		}

		close(out)
	}()

	return out
}

func main() {
	sum := 0
	for i := range generatePrimes(7919) {
		sum += i
	}
	fmt.Println(sum)
}
