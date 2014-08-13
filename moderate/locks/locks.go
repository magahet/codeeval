package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func shortcut(n, m int) int {
	if n <= 1 {
		return n - 1
	}
	c := n/2
	if m%2 == 1 {
		c += n/6	
	}
	if n%6 == 0 || (n%3 == 0 && m%2 == 1) {
		c++
	}
	return c
}

func naive(n, m int) int {
	l := make([]bool, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j += 2 {
			l[j] = true
		}
		for j := 2; j < n; j += 3 {
			l[j] = !l[j]
		}
	}
	l[len(l)-1] = !l[len(l)-1]
	var c int
	for j := 0; j < n; j++ {
		if !l[j] {
			c++
		}
	}
	return c
}

func processLine(line string) int {
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	
	return naive(n, m)
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
