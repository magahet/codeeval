package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func getSubStrings(s string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for l := len(s) / 2; l > 0; l-- {
			seen := make(map[string]bool)
			for i := 0; i < len(s)-l; i++ {
				if strings.Count(s[i:i+l], " ") == l {
					continue
				}
				if _, exists := seen[s[i:i+l]]; !exists {
					out <- s[i : i+l]
					seen[s[i:i+l]] = true
				}
			}
		}
	}()
	return out
}

func processLine(line string) string {
	for s := range getSubStrings(line) {
		if strings.Count(line, s) > 1 {
			return s
		}
	}

	return "NONE"
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
