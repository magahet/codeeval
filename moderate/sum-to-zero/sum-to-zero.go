package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
)

func combinations(iterable []int, r int) <-chan []int {
	pool := iterable
	n := len(pool)

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]int, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		c <- result
		var m, maxVal int
		for i := 1; i < int(new(big.Int).Binomial(int64(n), int64(r)).Int64()); i++ {
			m = r - 1
			maxVal = n - 1
			for indices[m] == maxVal {
				m--
				maxVal--
			}
			indices[m]++
			for j := m + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}
			result := make([]int, r)
			for ii, el := range indices {
				result[ii] = pool[el]
			}
			c <- result
		}
	}(c)

	return c

}

func sum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func processLine(line string) (count int) {
	numStrs := strings.Split(line, ",")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}
	//fmt.Println(nums)
	for set := range combinations(nums, 4) {
		//fmt.Println(set)
		if sum(set) == 0 {
			count++
		}
	}
	return
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
