package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"container/list"
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
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}


    /*
	for line := range readLine(file) {
	    fmt.Println(line)
	}
	*/
    num := 0
    lowest := 0
    lines := list.New()
	for line := range readLine(file) {
	    if num == 0 {
	        num, _ = strconv.Atoi(line)
	        continue
	    }
	    if lines.Len() == 0 {
	        lines.PushFront(line)
	        lowest = len(line)
	    } else if len(line) > lowest || lines.Len() < num {
	        for e := lines.Front(); e != nil; e = e.Next() {
	            if len(line) > len(e.Value.(string)) {
	                lines.InsertBefore(line, e)
	                break
	            }
	        }
	        if lines.Len() > num {
	            lines.Remove(lines.Back())
	        }
	    }
	}
	for e := lines.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}