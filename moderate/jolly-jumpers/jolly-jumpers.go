package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func processLine(line string) string {
	parts := strings.Fields(line)
	seen := make(map[int]bool)
	last := 0
	cur := 0

	for i, num := range parts[1:] {
		if i == 0 {
			last, _ = strconv.Atoi(num)
			continue
		}
		cur, _ = strconv.Atoi(num)
		seen[Abs(cur-last)] = true
		last = cur
	}

	if len(seen) != len(parts)-2 {
		return "Not jolly"
	}

	for i := 1; i < len(parts)-1; i++ {
		_, exists := seen[i]
		if exists == false {
			return "Not jolly"
		}
	}

	return "Jolly"
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
