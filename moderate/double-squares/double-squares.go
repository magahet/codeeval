package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func GCD(a, b uint32) uint32 {
    if a == 0 {
            return b
    }
    for b != 0 {
            if a > b {
                    a -= b
            } else {
                    b -= a
            }
    }
    return a
}

func findPrimeFactor(N uint32) uint32 {
	fmt.Println(N)
	if N == 1 {
		return 1
	} else if N%2 == 0 {
		return 2
	}
	x := rand.Uint32()%(N-2) + 1
	y := x
	c := rand.Uint32()%(N-2) + 1
	g := uint32(1)
	for g == 1 {
		x = ((x*x)%N + c) % N
		y = ((y*y)%N + c) % N
		y = ((y*y)%N + c) % N
        g = GCD(Abs(x-y), N)
		//fmt.Printf("gcd: %v, %v, %v, %v\n", x, y, big.NewInt(N), g)
	}
	//fmt.Printf("g:%v ", g)
	return g
}

func r2(current uint32) int {
	if current == 0 || current == 1 {
		return 1
	}

	a0 := 0
	b := 0

	for current > 1 {
		fmt.Println(current)
		prime := findPrimeFactor(current)
		for i := 0; big.NewInt(prime).ProbablyPrime(4) == false; i++ {
			prime = findPrimeFactor(current)
			if i > 5 {
				panic("could not find prime factor")
			}
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

	//fmt.Printf("b")
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

		X, _ := strconv.ParseUint(line, 10, 32)
		fmt.Printf("%d, %d\n", X, r2(X))
	}
}
