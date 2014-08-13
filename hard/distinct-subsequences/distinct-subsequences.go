package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func processLine(line string) int {
	parts := strings.Split(line, ",")
	chars := make([]string, len(parts[1]))
	for i, char := range parts[1] {
		chars[i] = string(char)
	}
	r := regexp.MustCompile(strings.Join(chars, `(.{0,1})`))
	fmt.Println(r.FindAllStringSubmatch(parts[0], -1))
	return len(r.FindAllStringSubmatch(parts[0], -1))
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
