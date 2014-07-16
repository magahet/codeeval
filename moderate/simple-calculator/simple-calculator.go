package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

def _parse(expression):
    tokens = []
    digits = ''
    if expression.count('(') != expression.count(')'):
        raise Exception('Unequal number of parentheses')
    for index, character in enumerate(expression):
        if character.isdigit() or character == '.':
            digits = digits + character
            if index < len(expression) - 1:
                continue
        if digits:
            tokens.append(float(digits))
        if character in '()^*/+-':
            tokens.append(character)
        digits = ''
    if '(' in tokens:
        return _group_tokens(tokens)
    else:
        return tokens

func processLine(line string) string {
	for i, char := range line {
	}
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
