package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func isUniform(s string) bool {
	return numSeq(s) == 1
}

func numSeq(s string) int {
	count := 1
	r := s[0]
	for i := 1; i < len(s); i++ {
		if r != s[i] {
			count++
		}
		r = s[i]
	}
	return count
}

func numSeqM(s string, m byte) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	r := s[0]
	if r == m {
		count++
	}
	for i := 1; i < len(s); i++ {
		if r != s[i] && r == m {
			count++
		}
		r = s[i]
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func processLine(b, p string, d int) string {
	if numSeqM(b, byte('1')) < numSeqM(p, byte('B')) {
		return "No"
	}
	if len(b) == 0 && len(p) > 0 {
		return "No"
	}
	if len(b) == 0 && len(p) == 0 {
		return "Yes"
	}
	if p[0] == 'B' && b[0] == '0' {
		return "No"
	}
	if p[len(p)-1] == 'B' && b[len(b)-1] == '0' {
		return "No"
	}
	if isUniform(p) && isUniform(b) {
		return "Yes"
	}

	i := 1
	if len(b) > 1 && len(p) > 1 && b[:2] == "10" && p[:2] == "BB" {
		i = strings.IndexRune(p, 'A')
	}
	for ; i < min(len(p), len(b)); i++ {
		//fmt.Printf("%d,", i)
		if processLine(b[1:], p[i:], d+1) == "Yes" {
			return "Yes"
		}
		if p[i] != p[0] {
			break
		}
	}
	return "No"
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
			parts := strings.Fields(line)
			fmt.Println(line)
			fmt.Println(processLine(parts[0], parts[1], 1))
		}
	}
}
