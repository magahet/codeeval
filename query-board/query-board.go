package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
)

type matrix struct {
    n, m int
    matrix []int
}

func newMatrix(n, m int) *matrix {
    matrix := new(matrix)
    matrix.n = n
    matrix.m = m
    matrix.matrix = make([]int, n * m)
    return matrix
}

func (m *matrix) setRow (x, v int) {
    for y := 0; y < m.m; y++ {
        m.matrix[x*m.m+y] = v
    }
}

func (m *matrix) setCol (y, v int) {
    for x := 0; x < m.n; x++ {
        m.matrix[x*m.m+y] = v
    }
}

func (m *matrix) queryRow (x int) int {
    sum := 0
    for y := 0; y < m.m; y++ {
        sum += m.matrix[x*m.m+y]
    }
    return sum
}

func (m *matrix) queryCol (y int) int {
    sum := 0
    for x := 0; x < m.n; x++ {
        sum += m.matrix[x*m.m+y]
    }
    return sum
}

func (m *matrix) processCmd(line string) {
    tokens := strings.Fields(line)
    cmd := tokens[0]
    rowOrCol, _ := strconv.Atoi(tokens[1])
    val := 0
    if len(tokens) == 3 {
        val, _ = strconv.Atoi(tokens[2])
    }
    switch cmd {
        default: fmt.Println(line)
        case "SetRow": m.setRow(rowOrCol, val)
        case "SetCol": m.setCol(rowOrCol, val)
        case "QueryRow": fmt.Println(m.queryRow(rowOrCol))
        case "QueryCol": fmt.Println(m.queryCol(rowOrCol))
    }
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
    m := newMatrix(256, 256)
	for line := range readLine(file) {
		m.processCmd(line)
	}
}
