package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type AgeRange struct {
	min, max int
	Name     string
}

func (ageRange *AgeRange) Contains(age int) bool {
	return age <= ageRange.max && age >= ageRange.min
}

var AgeGroups = []*AgeRange{
	&AgeRange{0, 2, "Home"},
	&AgeRange{3, 4, "Preschool"},
	&AgeRange{5, 11, "Elementary school"},
	&AgeRange{12, 14, "Middle school"},
	&AgeRange{15, 18, "High school"},
	&AgeRange{19, 22, "College"},
	&AgeRange{23, 65, "Work"},
	&AgeRange{66, 100, "Retirement"},
	&AgeRange{1 >> 31, -1, "This program is for humans"},
	&AgeRange{101, 1 << 31, "This program is for humans"},
}

func processLine(line string) string {
	age, _ := strconv.Atoi(line)
	for _, ageGroup := range AgeGroups {
		if ageGroup.Contains(age) {
			return ageGroup.Name
		}
	}
	return "None"
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

func parseLine(line string) (a int, b int, N int) {
	tokens := strings.Fields(line)
	a, _ = strconv.Atoi(tokens[0])
	b, _ = strconv.Atoi(tokens[1])
	N, _ = strconv.Atoi(tokens[2])
	return
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
