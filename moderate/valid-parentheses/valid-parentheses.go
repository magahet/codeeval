package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

var match = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func testString(line string) bool {
	if line == "" {
		return true
	} else if len(line)%2 != 0 {
		return false
	}

	depth := 1
	top := rune(line[0])

	for i, char := range line {
		if i == 0 {
			continue
		}
		if char == top {
			depth++
		} else if char == match[top] {
			depth--
		}
		if depth == 0 {
			if len(line) == 2 {
				return true
			} else if i < len(line)-1 {
				if testString(line[i+1:]) == false {
					return false
				}
			}

			return testString(line[1:i])

		} else if depth < 0 {
			return false
		}
	}

	if depth > 0 {
		return false
	}

	return true
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
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}
		}
	}
}
