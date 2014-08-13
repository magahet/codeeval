package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

var romanMap map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func processLine(line string) int {
	var pR, pRa, a, r, sum int
	for i := 0; i < len(line); i += 2 {
		a, _ = strconv.Atoi(string(line[i]))
		r = romanMap[rune(line[i+1])]
		if pR != 0 && r > pR {
			sum -= 2*pRa
		}
		pRa = a*r
		sum += pRa
		pR = r
	}
	return sum
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
