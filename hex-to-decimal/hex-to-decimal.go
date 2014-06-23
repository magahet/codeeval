package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"math"
)

func hexToDec(line string) int {
    sum := 0.0
    for i, char := range line {
        r := strings.NewReplacer("a", "10", "b", "11", "c", "12", "d", "13", "e", "14", "f", "15")
        hexDigit, _ := strconv.Atoi(r.Replace(string(char)))
        sum += float64(hexDigit) * math.Pow(16, float64(len(line) - 1 - i))
    }
    return int(sum)
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
	    fmt.Println(hexToDec(line))
	}
}
