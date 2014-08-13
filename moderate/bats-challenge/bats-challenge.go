package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)


func processLine(line string) int {
	parts := strings.Fields(line)
	length, _ := strconv.Atoi(parts[0])
	d, _ := strconv.Atoi(parts[1])
	n, _ := strconv.Atoi(parts[2])
	
	p := make([]int, n)
	var r int
	if n > 0 {
		for i, s := range parts[3:] {
			p[i], _ = strconv.Atoi(s)
		}
		r = p[0]
	} else {
		r = length + d
	}
	
	var c, j int
	l := -d
	for i := 6; i <= length - 6; i++ {
		if i <= r - d && i >= l + d {
			c++
			l = i
		} else if i >= r {
			l = r
			if j < len(p) - 1 {
				j++
				r = p[j]
			} else {
				r = length + d
			}
		}
	}
	
	return c
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
