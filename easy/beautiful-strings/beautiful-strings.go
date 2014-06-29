package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"
)

func processLine(line string) int {
	line = strings.ToLower(line)
	counts := make(map[rune]int)

	for _, char := range line {
		if unicode.IsLetter(char) {
			counts[char]++
		}
	}

	countArr := make([]int, len(counts))

	i := 0
	for _, count := range counts {
		countArr[i] = count
		i++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(countArr)))

	sum := 0
	for i, n := range countArr {
		sum += (26 - i) * n
	}

	return sum
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
