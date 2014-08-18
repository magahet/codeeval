package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
)

type Matrix struct {
	N int
	Value []int
}

func (a Matrix) Get(i, j int) int {
	if i < 0 || i >= a.N || j < 0 || j >= a.N {
		return 9999
	}
	return a.Value[i*a.N+j]
}

func (a Matrix) SmallestPathCost(i, j int) int {
	if i == j && i == a.N - 1 {
		return a.Get(i, j)
	} else if i == a.N - 1 {
		return a.SmallestPathCost(i, j+1) + a.Get(i, j)
	} else if j == a.N - 1 {
		return a.SmallestPathCost(i+1, j) + a.Get(i, j)
	} else {
		return min(a.SmallestPathCost(i+1, j), a.SmallestPathCost(i, j+1)) + a.Get(i, j)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findSum(n int, line string) int {
	v := make([]int, n*n)
	for i, str := range strings.Split(line, ",") {
		v[i], _ = strconv.Atoi(str)
	}
	a := Matrix{n, v}
	
	return a.SmallestPathCost(0, 0)
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
	
	rows := ""
	var n int
	for line := range readLine(file) {
		if !strings.Contains(line, ",") {
			if len(rows) > 0 {
				fmt.Println(findSum(n, strings.Trim(rows, ",")))
			}
			n, _ = strconv.Atoi(line)
			rows = ""
		} else {
			rows += line + ","
		}
	}
	fmt.Println(findSum(n, strings.Trim(rows, ",")))
}
