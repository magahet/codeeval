package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var currency = map[int]string{
	10000: "ONE HUNDRED",
	5000:  "FIFTY",
	2000:  "TWENTY",
	1000:  "TEN",
	500:   "FIVE",
	200:   "TWO",
	100:   "ONE",
	50:    "HALF DOLLAR",
	25:    "QUARTER",
	10:    "DIME",
	5:     "NICKEL",
	1:     "PENNY",
}

var order = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 25, 10, 5, 1}

func parseAmount(s string) int {
	if strings.ContainsRune(s, '.') == false {
		s += "00"
	}

	s = strings.Replace(s, ".", "", 1)
	a, _ := strconv.Atoi(s)
	return a

}

func processLine(line string) string {
	parts := strings.Split(line, ";")
	pp := parseAmount(parts[0])
	ch := parseAmount(parts[1])
	change := ch - pp
	//fmt.Println(pp, ch, change)

	if change < 0 {
		return "ERROR"
	} else if change == 0 {
		return "ZERO"
	}

	results := ""
	for _, value := range order {
		for change >= value {
			results += fmt.Sprintf("%s,", currency[value])
			change -= value
		}
	}

	return strings.Trim(results, ",")
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
