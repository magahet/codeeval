package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

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
		fmt.Println("error opening file", os.Args[1], ":", err) os.Exit(1)
	}

    // prepopulate set of perfect squares up to limit
    perfectSquares := make(map[int]bool)
	for i := 0; i <= 32767; i++ {
	    perfectSquares[i*i] = true
	}

    var N int
	for line := range readLine(file) {
	    if N == 0 {
	        N, _ = strconv.Atoi(line)
	        continue
	    }

	    X, _ := strconv.Atoi(line)
	    count := 0
	    // iterate int combinations
	    for i := 0; i <= X/2; i++ {
	        _, e1 := perfectSquares[i]
	        _, e2 := perfectSquares[X-i]
	        // test if both ints are in perfect squares set
	        if e1 && e2 {
	            count++
	        }
	    }
	    fmt.Println(count)
	}
}
