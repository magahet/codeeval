package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) int {
	parts := strings.Split(line, ",")
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])
	a, _ := strconv.Atoi(parts[2])
	b, _ := strconv.Atoi(parts[3])
	c, _ := strconv.Atoi(parts[4])
	r, _ := strconv.Atoi(parts[5])

	seenMap := make(map[int]bool)
	seenMap[a] = true
	seenLog := make(chan int, k)
	defer close(seenLog)
	seenLog <- a

	m := a
	for i := 1; i < k; i++ {
		m = (b*m + c) % r
		seenMap[m] = true
		seenLog <- m
	}

	var l int
	for i := k; i < n; i++ {
		m = findMin(seenMap)
		seenMap[m] = true
		l = <-seenLog
		seenLog <- m
		delete(seenMap, l)
	}

	return m
}

func findMin(seen map[int]bool) int {
	for i := 0; i <= len(seen); i++ {
		if _, exists := seen[i]; !exists {
			return i
		}
	}
	return len(seen)
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
			//processLine(line)
		}
	}
}
