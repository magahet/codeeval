package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func Count(line, sub string) int {
	if len(sub) == 0 {
		return 1
	}
	char := rune(sub[0])	
	count := 0
	for i, c := range line {
		if c == char {
			count += Count(line[i+1:], sub[1:])
		}
	}	
	return count
}

func processLine(line string) int {
	parts := strings.Split(line, ",")
	return Count(parts[0], parts[1])
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
