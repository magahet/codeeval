package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

var ZERO *big.Int = big.NewInt(0)
var ONE *big.Int = big.NewInt(1)
var TWO *big.Int = big.NewInt(2)
var THREE *big.Int = big.NewInt(3)
var FOUR *big.Int = big.NewInt(4)

// findPrimeFactor implements Pollard's Rho algorithm
func findPrimeFactor(N *big.Int) *big.Int {
	// N == 1
	if N.Cmp(ONE) == 0 {
		return big.NewInt(1)
		// N is even
	} else if big.NewInt(0).Mod(N, TWO).Cmp(ZERO) == 0 {
		return big.NewInt(2)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	max := big.NewInt(0).Sub(N, TWO)
	x := big.NewInt(0)
	x.Add(x.Rand(r, max), ONE)
	y := big.NewInt(0).Set(x)
	c := big.NewInt(0)
	c.Add(c.Rand(r, max), ONE)
	g := big.NewInt(1)
	d := big.NewInt(0)
	// g == 1
	for g.Cmp(ONE) == 0 {
		x.Mul(x, x).Mod(x, N).Add(x, c).Mod(x, N)
		y.Mul(y, y).Mod(y, N).Add(y, c).Mod(y, N)
		y.Mul(y, y).Mod(y, N).Add(y, c).Mod(y, N)
		g.GCD(nil, nil, d.Abs(d.Sub(x, y)), N)
	}
	if g.Cmp(ZERO) == 0 {
		return N
	}

	return g
}

// r2 implements Beiler's algorithm for counting unique sum of squares sets
func r2(current *big.Int) int {
	if current.Cmp(ZERO) == 0 || current.Cmp(ONE) == 0 {
		return 1
	}

	a0 := 0
	b := 0
    sq := big.NewInt(0)
    zero := 0

	// current != 1
	for current.Cmp(ONE) != 0 {
		prime := findPrimeFactor(current)
		isPrime := prime.ProbablyPrime(4)
		for i := 0; isPrime == false; i++ {
			prime = findPrimeFactor(current)
			if i > 10 {
				panic("could not find a prime factor")
			}
			isPrime = prime.ProbablyPrime(4)
		}

        if sq.Cmp(ZERO) == 0 {
            if sq.Mul(prime, prime).Cmp(current) == 0 {
                zero = 1
            }
        }

		count := 0
		i := big.NewInt(0)
		for i.Mod(current, prime).Cmp(ZERO) == 0 {
			count++
			current.Div(current, prime)
			if prime.Cmp(ONE) == 0 {
				break
			}
		}

		bluePrime := big.NewInt(0)
		redPrime := big.NewInt(0)

		if prime.Cmp(TWO) == 0 {
			a0 += count
		} else if bluePrime.Sub(prime, THREE).Mod(bluePrime, FOUR).Cmp(ZERO) == 0 && count%2 != 0 {
			return 0
		} else if redPrime.Sub(prime, ONE).Mod(redPrime, FOUR).Cmp(ZERO) == 0 {
			if b == 0 {
				b = count + 1
			} else {
				b *= count + 1
			}
		}
	}

	if b%2 == 0 {
		return (b / 2) + zero
	} else if a0%2 == 0 {
		return ((b - 1) / 2) + zero
	}

	return ((b + 1) / 2) + zero

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
	X := new(big.Int)
	for line := range readLine(file) {
		if N == 0 {
			N, _ = strconv.Atoi(line)
			continue
		}

		X.SetString(line, 10)
		fmt.Println(r2(X))
	}
}
