package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) string {
	setStrs := strings.Split(line, ";")
	set1 := strings.Split(setStrs[0], ",")
	set2 := strings.Split(setStrs[1], ",")

	//m1, _ := strconv.Atoi(string(set1[0]))
	m2, _ := strconv.Atoi(string(set1[len(set1)-1]))
	m3, _ := strconv.Atoi(string(set2[0]))
	m4, _ := strconv.Atoi(string(set2[len(set2)-1]))

	if m2 < m3 {
		return ""
	}

	start := 0
	for i, num := range set1 {
		if num == set2[0] {
			start = i
		}
	}

	result := ""
	for i := start; i < len(set1); i++ {
		m, _ := strconv.Atoi(string(set1[1]))
		if m > m4 {
			break
		}
		result += set1[i] + ","
	}

	return strings.Trim(result, ",")
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
		fmt.Println(processLine(line))
	}
}
