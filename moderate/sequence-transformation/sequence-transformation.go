package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func processLine(line string) string {
	parts := strings.Fields(line)
	rStr := "^"
	for _, char := range parts[0] {
		if char == '0' {
			rStr += "A+"
		} else {
			rStr += `(A+|B+)`
		}
	}
	r, _ := regexp.Compile(rStr + "$")
	if r.MatchString(parts[1]) {
		return "Yes"
	}
	return "No"
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
			//fmt.Println(line)
			fmt.Println(processLine(line))
		}
	}
}
