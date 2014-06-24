package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func toRoman(line string) string {
	n, _ := strconv.Atoi(line)
	result := ""
	for n >= 1000 {
		n -= 1000
		result += "M"
	}
	if n >= 900 {
		n -= 900
		result += "CM"
	}
	for n >= 500 {
		n -= 500
		result += "D"
	}
	if n >= 400 {
		n -= 400
		result += "CD"
	}
	for n >= 100 {
		n -= 100
		result += "C"
	}
	if n >= 90 {
		n -= 90
		result += "XC"
	}
	if n >= 50 {
		n -= 50
		result += "L"
	}
	if n >= 40 {
		n -= 40
		result += "XL"
	}
	for n >= 10 {
		n -= 10
		result += "X"
	}
	if n == 9 {
		n -= 9
		result += "IX"
	}
	for n >= 5 {
		n -= 5
		result += "V"
	}
	if n == 4 {
		n -= 4
		result += "IV"
	}
	for n >= 1 {
		n -= 1
		result += "I"
	}
	return result
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
		fmt.Println(toRoman(line))
	}
}
