package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
)

func processLine(line string) string {
	n, _ := strconv.Atoi(line)
	sum := 0.0
	for _, digitStr := range line {
		digit, _ := strconv.Atoi(string(digitStr))
		sum += math.Pow(float64(digit), float64(len(line)))
	}
	if float64(n) == sum {
		return "True"
	} else {
		return "False"
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

	for line := range readLine(file) {
		fmt.Println(processLine(line))
	}
}
