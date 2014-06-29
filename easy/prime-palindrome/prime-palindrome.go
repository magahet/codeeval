package main

import (
	"fmt"
	"math"
	"strconv"
)

func largestPrimePalindrome(N int) int {
	largest := 0
	notPrime := make([]bool, N+1)
	limit := int(math.Sqrt(float64(N)))
	for i := 2; i <= limit; i++ {
		if notPrime[i] == false {
			if i > largest && isPalindrome(i) {
				largest = i
			}
			for j := i * i; j <= N; j += i {
				notPrime[j] = true
			}
		}
	}

	for i := limit + 1; i <= N; i++ {
		if notPrime[i] == false {
			if i > largest && isPalindrome(i) {
				largest = i
			}
		}
	}

	return largest
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(largestPrimePalindrome(1000))
}
