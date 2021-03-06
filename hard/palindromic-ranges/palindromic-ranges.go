package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func getSubRanges(s, e int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		for i := s; i <= e; i++ {
			for j := i; j <= e; j++ {
				subRange := make([]int, j-i+1)
				for k := 0; k <= j-i; k++ {
					subRange[k] = i + k
				}
				out <- subRange
			}
		}
	}()
	return out
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func countPalindromes(subRange []int) int {
	count := 0
	for _, i := range subRange {
		if isPalindrome(i) {
			count++
		}
	}
	return count
}

func processLine(line string) int {
	parts := strings.Fields(line)
	s, _ := strconv.Atoi(parts[0])
	e, _ := strconv.Atoi(parts[1])
	//fmt.Println(s, e)
	count := 0
	for subRange := range getSubRanges(s, e) {
		if countPalindromes(subRange)%2 == 0 {
			count++
		}
	}
	return count
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
