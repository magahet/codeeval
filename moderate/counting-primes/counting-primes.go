package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func genPrimes(N int) []bool {
	notPrime := make([]bool, N+1)
	limit := int(math.Sqrt(float64(N)))
	for i := 2; i <= limit; i++ {
		if notPrime[i] == false {
			for j := i * i; j <= N; j += i {
				notPrime[j] = true
			}
		}
	}
	return notPrime
}

func processLine(line string, primes []bool) int {
	parts := strings.Split(line, ",")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	count := 0
	for i := start; i <= end; i++ {
		if primes[i] == false {
			count++
		}
	}
	return count
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

	primes := genPrimes(500)
	for line := range readLine(file) {
		if line != "" {
			fmt.Println(processLine(line, primes))
		}
	}
}
