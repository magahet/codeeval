package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

var A Matrix = Matrix{3, 4, "ABCESFCSADEE"}

type Matrix struct {
	M, N  int
	Value string
}

func (a Matrix) Get(i, j int) rune {
	if i < 0 || i >= a.M || j < 0 || j >= a.N {
		return ' '
	}
	return rune(a.Value[i*a.N+j])
}

func (a Matrix) Index(i, j int) int {
	return i*a.N + j
}

func (a Matrix) String() string {
	results := ""
	for i := 0; i < a.M; i++ {
		results += a.Value[i*a.N:i*a.N+a.N] + "\n"
	}
	return results
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (a Matrix) Contains(s string, p int, path []int) bool {
	if len(s) == 0 {
		return true
	}
	firstChar := rune(s[0])
	newPath := append(path, p)
	if p == -1 {
		for i, char := range a.Value {
			if char == firstChar && a.Contains(s[1:], i, newPath) {
				return true
			}
		}
	} else {
		i := p / a.N
		j := p % a.N
		for _, d := range []int{-1, 1} {
			if !contains(newPath, a.Index(i+d, j)) && a.Get(i+d, j) == firstChar && a.Contains(s[1:], a.Index(i+d, j), newPath) {
				return true
			}
			if !contains(newPath, a.Index(i, j+d)) && a.Get(i, j+d) == firstChar && a.Contains(s[1:], a.Index(i, j+d), newPath) {
				return true
			}
		}
	}
	return false
}

func processLine(line string) bool {
	return A.Contains(line, -1, []int{})
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
			if processLine(line) {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}
		}
	}
}
