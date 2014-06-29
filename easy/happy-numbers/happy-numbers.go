package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func isHappy(line string) int {
	sumMap := make(map[int]bool)
	n, _ := strconv.Atoi(line)
	for j := 0; j < 10; j++ {
		if n == 1 {
			return 1
		} else {
			nStr := strconv.Itoa(n)
			n = 0
			for _, char := range nStr {
				i, _ := strconv.Atoi(string(char))
				n += i * i
			}
			if sumMap[n] == true {
				return 0
			} else {
				sumMap[n] = true
			}
		}
	}
	return 0
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
		fmt.Println(isHappy(line))
	}
}
