package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

const valid = "abcdefghijklmnopqrstuvwxyz :()"

func testString(line string) bool {
	var depth int
	for i, char := range line {
		if !strings.ContainsRune(valid, char) {
			return false
		}
		if depth < 0 && strings.Count(line[:i], ":)")+depth < 0 {
			return false
		}
		if char == '(' {
			depth++
		} else if char == ')' {
			depth--
		}
	}

	var s rune
	var d int
	if depth > 0 {
		s = '('
		d = -1
	} else if depth < 0 {
		s = ')'
		d = 1
	}

	for i, char := range line {
		if depth == 0 {
			return true
		}
		if i == 0 {
			continue
		}
		if char == s && rune(line[i-1]) == ':' {
			depth += d
		}
	}

	return depth == 0
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
			if testString(line) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
