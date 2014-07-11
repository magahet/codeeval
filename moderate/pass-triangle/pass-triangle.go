package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func Cmp(set, set2 map[string]bool) bool {
	if len(set) != len(set2) {
		return false
	}
	for k := range set {
		_, ok := set2[k]
		if !ok {
			return false
		}
	}
	return true
}

func sqrt(n int) int {
    for i := 0; i*i <= n; i++ {
        if i*i == n {
            return i
        }
    }
    return -1
}

func processLine(line string) string {
	parts := strings.Split(line, ";")
	n, _ := strconv.Atoi(parts[0])
	nums := strings.Split(parts[1], ",")

	// create validation set
	valSet := make(map[string]bool)
	for i := 0; i < n; i++ {
		valSet[strconv.Itoa(i+1)] = true
	}

    var set map[string]bool
	// check rows
	for i := 0; i < n; i++ {
		set = make(map[string]bool)
		for j := 0; j < n; j++ {
			set[nums[i+j*n]] = true
		}
		if !Cmp(set, valSet) {
			return "False"
		}
	}

	// check cols
	for i := 0; i < n; i++ {
		set = make(map[string]bool)
		for j := 0; j < n; j++ {
			set[nums[i*n+j]] = true
		}
		if !Cmp(set, valSet) {
			return "False"
		}
	}

	// check boxes
	for i := 0; i < sqrt(n); i++ {
		for j := 0; j < sqrt(n); j++ {
    		set = make(map[string]bool)
		    for k := 0; k < sqrt(n); k++ {
    		    for l := 0; l < sqrt(n); l++ {
    		        //fmt.Printf("%d,", i*n*sqrt(n) + j*sqrt(n) + k*n + l)
    			    set[nums[i*n*sqrt(n) + j*sqrt(n) + k*n + l]] = true
    		    }
		    }
		    //fmt.Printf("%v\n", set)
    		if !Cmp(set, valSet) {
    		    //fmt.Println("%v", set)
    			return "False"
    		}
		}
	}

    return "True"
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
