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

func processLine(line string) int {
	duplicates := make(map[string]bool)
	for _, num := range strings.Fields(line) {
		_, exists := duplicates[num]
		duplicates[num] = exists
	}

	lowest := math.MaxInt32
	for num, dup := range duplicates {
		if dup == false {
			n, _ := strconv.Atoi(num)
			if n < lowest {
				lowest = n
			}
		}
	}

	lowestStr := strconv.Itoa(lowest)
	for i, num := range strings.Fields(line) {
		if num == lowestStr {
			return i + 1
		}
	}
	return 0
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
