package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"unicode/utf8"
	"strings"
)

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

func truncate(line string) string {
	length := utf8.RuneCountInString(line)
	if length <= 55 {
		return line
	} else {
		index := strings.LastIndex(line[:40], " ")
		if index == -1 {
			index = 40
		}
		return line[0:index] + "... <Read More>"
	}
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
		fmt.Println(truncate(line))
	}
}
