package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var r = strings.NewReplacer("0", "1", "1", "2", "2", "0")

func genPhrase(line string) string {
	return line + r.Replace(line)
}

func nearestPower2(n int) int {
	p := 1
	for p <= n {
		p *= 2
	}
	return p / 2
}

func processLine(b string, n int) string {
	n -= nearestPower2(n)
	if n == 0 {
		return r.Replace(b)
	}
	return processLine(r.Replace(b), n)
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
			n, _ := strconv.Atoi(line)
			if n == 0 {
				fmt.Println("0")
			} else {
				fmt.Println(processLine("0", n))
			}
		}
	}
}
