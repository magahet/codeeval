package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Matrix struct {
	M, N int
	Value string
}

func (a Matrix) Get(i, j int) rune {
	if i < 0 || i >= a.M || j < 0 || j >= a.N {
		return ' '
	}
	return rune(a.Value[i*a.N+j])
}

func (a Matrix) String() string {
	results := ""
	for i := 0; i < a.M; i++ {
		results += a.Value[i*a.N:i*a.N+a.N] + "\n"
	}
	return results
}

func (a Matrix) CountAdjMines(i, j int) int {
	c := 0
	
	//fmt.Println(i, j)
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if a.Get(i+di, j+dj) == '*' {
				c++
				//fmt.Println(i+di, j+dj, c)
			}
		}
	}
	
	return c
}

func printR(line string, m, n int) string {
	results := ""
	for i := 0; i < m; i++ {
		results += line[i*n:i*n+n] + "\n"
	}
	return results
}

func processLine(line string) string {
	parts := strings.Split(line, ";")
	mn := strings.Split(parts[0], ",")
	m, _ := strconv.Atoi(mn[0])
	n, _ := strconv.Atoi(mn[1])
	a := Matrix{m, n, parts[1]}
	
	result := ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a.Get(i, j) == '*' {
				result += "*"
			} else {
				result += strconv.Itoa(a.CountAdjMines(i, j))
			}
		}
	}
	return result
	//return printR(result, m, n)
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
