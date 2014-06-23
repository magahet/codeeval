package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func findRep(line string) int {
    pattern := ""
    for _, char := range line {
        pattern += string(char)
        if len(pattern) == len(line) {
            return len(line)
        } else if len(line) % len(pattern) == 0 {
            if strings.Count(line, pattern) == len(line) / len(pattern) {
                return len(pattern)
            }
        }
    }
    return len(pattern)
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
		fmt.Println(findRep(line))
	}
}
