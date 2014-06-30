package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func findRep(line string) string {
    chain := make(map[string]string)
    last := ""
    result := ""
    for _, num := range strings.Fields(line) {
        if last != "" {
            chain[last] = num
        }
        _, ok := chain[num]
        if ok {
            //fmt.Printf("%v\n", chain)
            result += num + " "
            for cur := chain[num]; cur != num; cur = chain[cur] {
                result += cur + " "
            }
            break
        }
        last = num
    }
    return strings.Trim(result, " ")
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
		fmt.Println(findRep(line))
	}
}
