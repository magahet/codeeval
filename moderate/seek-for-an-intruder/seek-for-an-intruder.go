package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func binToDec(s string) string {
	i, _ := strconv.ParseUint(s, 2, 64)
	return fmt.Sprintf("%d", i)
}

func octOrHexToDec(s string) string {
	i, _ := strconv.ParseUint(s, 0, 64)
	return fmt.Sprintf("%d", i)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	content := string(file)

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	ipCount := make(map[string]int)

	r, _ := regexp.Compile(`[^0-9a-fA-F.]`)
	content = r.ReplaceAllString(content, " ")
	//r, _ = regexp.Compile(`(^|\s)\S{1,2}(\s|$)`)
	//content = r.ReplaceAllString(content, " ")
	//fmt.Println(content)
	//fmt.Println("--------------")
	for _, str := range strings.Fields(content) {
		dotCount := strings.Count(str, ".")
		if len(str) < 7 || dotCount > 0 && dotCount != 3 {
			continue
		}
		fmt.Println(str)
	}

	//r, _ = regexp.Compile(`(0|1){8,32}`)
	//content = r.ReplaceAllStringFunc(content, binToDec)

	//r, _ = regexp.Compile(`0\d{1,11}`)
	//content = r.ReplaceAllStringFunc(content, octOrHexToDec)

	//r, _ = regexp.Compile(`0x[0-9a-fA-F]{2,8}`)
	//content = r.ReplaceAllStringFunc(content, octOrHexToDec)
	//fmt.Println(content)

	// dec
	r, _ = regexp.Compile(`\d{10}`)
	for _, s := range r.FindAllString(content, -1) {
		a, _ := strconv.ParseUint(s, 10, 64)
		ipStr := fmt.Sprintf("%d.%d.%d.%d", byte(a>>24), byte(a>>16), byte(a>>8), byte(a))
		//fmt.Println(s, a)
		ip := net.ParseIP(ipStr)
		if ip != nil {
			ipCount[ip.String()]++
		}
	}

	// dotted dec
	r, _ = regexp.Compile(`(?:\d{1,3}\.){3}\d{1,3}`)
	for _, s := range r.FindAllString(content, -1) {
		ip := net.ParseIP(s)
		if ip != nil {
			ipCount[ip.String()]++
		}
	}
	fmt.Println(ipCount)

}
