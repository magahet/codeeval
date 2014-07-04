package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"path"
	"strconv"
)

func findPrimeFactor(N int64) int {
	if N == 1 {
		return 1
	} else if N%2 == 0 {
		return 2
	}
	x := int64(rand.Int63n(N-2) + 1)
	y := x
	c := int64(rand.Int63n(N-2) + 1)
	g := int64(1)
	for g == 1 {
		fmt.Print('.')
		x = ((x*x)%N + c) % N
		y = ((y*y)%N + c) % N
		y = ((y*y)%N + c) % N
		g = big.NewInt(0).GCD(nil, nil, big.NewInt(0).Sub(big.NewInt(x), big.NewInt(y)), big.NewInt(N)).Int64()
	}
	return int(g)
}

func compositeTest(a, d, n, s int64) bool {
	if int64(math.Pow(float64(a), float64(d)))%n == 1 {
		return false
	}
	for i := int64(0); i < s; i++ {
		if int64(math.Pow(float64(a), math.Pow(2.0, float64(i*d))))%n == n-1 {
			return false
		}
	}
	return true // n  is definitely composite
}

func primeTest(n int64) bool {
	if n >= 1 && n <= 3 {
		return true
	}

	d := n - 1
	s := int64(0)

	for d%2 != 0 {
		d = d >> 1
		s++
	}

	var a []int64
	if n < 1373653 {
		a = []int64{2, 3}
	} else if n < 25326001 {
		a = []int64{2, 3, 5}
	} else if n == 3215031751 {
		return false
	} else if n < 118670087467 {
		a = []int64{2, 3, 5, 7}
	}

	for _, a_i := range a {
		if compositeTest(a_i, d, n, s) {
			return false
		}
	}

	return true
}

func r2(current int) int {
	if current == 0 || current == 1 {
		return 1
	}

	a0 := 0
	b := 0

	for current > 1 {
		prime := findPrimeFactor(int64(current))
		if primeTest(int64(prime)) == false {
			continue
		}

		count := 0
		for current%prime == 0 {
			count++
			current /= prime
		}

		if prime == 2 {
			a0 += count
		} else if (prime-3)%4 == 0 && count%2 != 0 {
			return 0
		} else if (prime-1)%4 == 0 {
			if b == 0 {
				b = count + 1
			} else {
				b *= count + 1
			}
		}
	}

	if b%2 == 0 {
		return b / 2
	} else if a0%2 == 0 {
		return (b - 1) / 2
	} else {
		return (b + 1) / 2
	}

}

func readLine(file *os.File) <-chan string {
	out := make(chan string)
	go func() {
		in := bufio.NewReader(file)
		linePartial := ""
		for {
			bytes, isPrefix, err := in.ReadLine()
			if err != nil {
				break
			} else if isPrefix {
				linePartial += string(bytes)
			} else {
				out <- linePartial + string(bytes)
				linePartial = ""
			}
		}
		close(out)
	}()
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	var N int
	for line := range readLine(file) {
		if N == 0 {
			N, _ = strconv.Atoi(line)
			continue
		}

		X, _ := strconv.Atoi(line)
		fmt.Println(r2(X))
	}
}
