package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type intSlice []int

func (p intSlice) More (i, j int) bool {
    return p[i] > p[j]
}

func (p intSlice) Swap (i, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p intSlice) IsSorted () bool {
	for i := 0; i < len(p) - 1; i++ {
		if p.More(i, i+1) {
			return false
		}
	}
	return true
}

func (p intSlice) Bubble() {
	for i := 0; i < len(p) - 1; i++ {
		if p.More(i, i+1) {
			p.Swap(i, i+1)
		}
	}
}

func (p intSlice) String() string {
	s := make([]string, len(p))
	for i, n := range p {
		s[i] = strconv.Itoa(n)
	}
	return strings.Join(s, " ")
}
	

func processLine(line string) string {
    parts := strings.Split(line, " | ")
    nums := strings.Fields(parts[0])
    list := make(intSlice, len(nums))
    for i, n := range nums {
    	list[i], _ = strconv.Atoi(n)
    }
    iters, _ := strconv.Atoi(parts[1])
    for i := 0; i < iters; i++ {
    	if list.IsSorted() {
    		break
    	}
    	list.Bubble()
    }
    return list.String()
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
