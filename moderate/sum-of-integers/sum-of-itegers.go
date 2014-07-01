package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
)

func maxBlock(ints []int, blockSize int) int {
    var sum, max int
    max = -1<<31
    for i := 0; i < len(ints); i++ {
        sum += ints[i]
        if i >= blockSize {
            sum -= ints[i-blockSize]
        }
        if sum > max {
            max = sum
        }
    }
    return max
}

func processLine(line string) int {
    tokens := strings.Split(line, ",")
    ints := make([]int, len(tokens))
    for i, str := range tokens {
        ints[i], _ = strconv.Atoi(strings.Trim(str, " "))
    }
    var cur, max int
    max = -1<<31
    for i := 1; i <= len(ints); i++ {
        cur = maxBlock(ints, i)
        if cur > max {
            max = cur
        }
    }
    return max
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
