package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	rows := make([][]int, len(lines))
	for i, line := range lines {
		rows[i] = make([]int, i+1)
		for j, numStr := range strings.Fields(line) {
			rows[i][j], _ = strconv.Atoi(numStr)
		}
	}

	for i := len(lines) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			rows[i-1][j] += maxInt(rows[i][j], rows[i][j+1])
		}
		//fmt.Println(rows[i])
	}
	fmt.Println(rows[0][0])
}
