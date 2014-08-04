package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type IP []int

func parseIP(ipStr string) IP {
	p := make(IP, 4)
	for i, s := range strings.Split(ipStr, ".") {
		p[i], _ = strconv.Atoi(s)
		if p[i] < 0 || p[i] > 255 {
			return nil
		}
	}
	if p[0] == 255 && p[1] == 255 && p[2] == 255 && p[3] == 255 {
		return nil
	}
	if p[0] < 1 {
		return nil
	}
	return p
}

func (ip IP) String() string {
	octets := make([]string, 4)
	for i, d := range ip {
		octets[i] = strconv.Itoa(d)
	}
	return strings.Join(octets, ".")
}

func binToDec(s string) string {
	i, _ := strconv.ParseUint(s, 2, 64)
	return fmt.Sprintf("%d", i)
}

func dotBinToDec(s string) string {
    octets := make([]string, 4)
    for i, o := range strings.Split(s, ".") {
        octets[i] = binToDec(o)
    }
	return strings.Join(octets, ".")
}

func octOrHexToDec(s string) string {
	i, _ := strconv.ParseUint(s, 0, 64)
	return fmt.Sprintf("%d", i)
}

func dotOctOrHexToDec(s string) string {
    octets := make([]string, 4)
    for i, o := range strings.Split(s, ".") {
        octets[i] = octOrHexToDec(o)
    }
	return strings.Join(octets, ".")
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
	newContent := ""
	for _, str := range strings.Fields(content) {
		dotCount := strings.Count(str, ".")
		if len(str) < 7 || (dotCount > 0 && dotCount != 3) {
			continue
		}
		newContent += fmt.Sprintf("%s\n", str)
	}
	content = newContent

	r, _ = regexp.Compile(`[01]{32}`)
	content = r.ReplaceAllStringFunc(content, binToDec)

	r, _ = regexp.Compile(`[01]{8}\.[01]{8}\.[01]{8}\.[01]{8}`)
	content = r.ReplaceAllStringFunc(content, dotBinToDec)

	r, _ = regexp.Compile(`0\d{11}`)
	content = r.ReplaceAllStringFunc(content, octOrHexToDec)

	r, _ = regexp.Compile(`0\d{1,4}\.0\d{1,4}\.0\d{1,4}\.0\d{1,4}`)
	content = r.ReplaceAllStringFunc(content, dotOctOrHexToDec)

	//r, _ = regexp.Compile(`0x[0-9a-fA-F]{8}`)
	//content = r.ReplaceAllStringFunc(content, catHexToDec)

	r, _ = regexp.Compile(`0x[0-9a-fA-F]{1,2}\.0x[0-9a-fA-F]{1,2}\.0x[0-9a-fA-F]{1,2}\.0x[0-9a-fA-F]{1,2}`)
	content = r.ReplaceAllStringFunc(content, dotOctOrHexToDec)

	//fmt.Println(content)
	// dec
	r, _ = regexp.Compile(`\d{10,}`)
	for _, s := range r.FindAllString(content, -1) {
		a, _ := strconv.ParseUint(s, 10, 64)
		ipStr := fmt.Sprintf("%d.%d.%d.%d", byte(a>>24), byte(a>>16), byte(a>>8), byte(a))
		//fmt.Println(s, a)
		ip := parseIP(ipStr)
		if ip != nil {
			ipCount[ip.String()]++
		}
	}

	// dotted dec
	r, _ = regexp.Compile(`[12]*\d{1,2}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	for _, s := range r.FindAllString(content, -1) {
		ip := parseIP(s)
		if ip != nil {
			ipCount[ip.String()]++
		}
	}

    max := 0
    ip := ""
	for k, v := range ipCount {
	    if v > max {
	        max = v
	        ip = k
	    }
	}

	fmt.Println(ip)

}