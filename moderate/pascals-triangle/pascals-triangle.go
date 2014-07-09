package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) string {
	rows, _ := strconv.Atoi(line)
	result := ""
	c := big.NewInt(0)
	for k := int64(0); k < int64(rows); k++ {
		for n := int64(0); n <= k; n++ {
			result += fmt.Sprintf("%d ", c.Binomial(k, n))
		}
	}
	return strings.Trim(result, " ")
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
