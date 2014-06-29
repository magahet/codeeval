package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func processLine(line string) {
	words := ""
	numbers := ""
	for _, word := range strings.Split(line, ",") {
		if unicode.IsNumber(rune(word[0])) {
			numbers += word + ","
		} else {
			words += word + ","
		}
	}
	if len(words) > 0 {
		fmt.Print(strings.Trim(words, ","))
		if len(numbers) > 0 {
			fmt.Print("|")
		}
	}
	if len(numbers) > 0 {
		fmt.Print(strings.Trim(numbers, ","))
	}
	fmt.Println("")
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
		processLine(line)
	}
}
