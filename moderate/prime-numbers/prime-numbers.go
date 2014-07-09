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

func processLine(line string) string {
	N, _ := strconv.Atoi(line)
	notPrime := make([]bool, N+1)
	limit := int(math.Sqrt(float64(N)))
	for i := 2; i <= limit; i++ {
		if notPrime[i] == false {
			for j := i * i; j <= N; j += i {
				notPrime[j] = true
			}
		}
	}

	results := ""
	for i := 2; i < N; i++ {
		if notPrime[i] == false {
			results += fmt.Sprintf("%d,", i)
		}
	}

	return strings.Trim(results, ",")
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

	for line := range readLine(file) {
		if line != "" {
			fmt.Println(processLine(line))
		}
	}
}
