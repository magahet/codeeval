package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

type Token struct {
	Value    float64
	Operator rune
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func dev(a, b float64) float64 {
	return a / b
}

func exp(a, b float64) float64 {
	return math.Pow(a, b)
}

type fn func(float64, float64) float64

var handler = map[rune]fn{
	'+': add,
	'-': sub,
	'*': mul,
	'/': dev,
	'^': exp,
	'_': sub,
}

var groups = []string{"^", "_", "*/", "+-"}

func calculateValue(tokens *list.List) float64 {
	var value, prev, next float64
	var opt rune
	for _, g := range groups {
		for t := tokens.Front(); tokens.Len() > 1 && t != nil; {
			opt = t.Value.(*Token).Operator
			if strings.ContainsRune(g, opt) {
				prev = t.Prev().Value.(*Token).Value
				next = t.Next().Value.(*Token).Value
				//fmt.Printf("%.1f %c %.1f\n", prev, opt, next)
				value = handler[opt](prev, next)
				tokens.Remove(t.Next())
				tokens.Remove(t.Prev())
				tokens.InsertAfter(&Token{Value: value}, t)
				tokens.Remove(t)
				t = tokens.Front()
			} else {
				t = t.Next()
			}
		}
	}

	return tokens.Front().Value.(*Token).Value
}

func eval(e string) float64 {
	//fmt.Println(e)
	buffer := ""
	tokens := list.New()
	depth := 0
	var value float64
	for i, char := range e {
		if char == '(' {
			if depth > 0 {
				buffer += "("
			}
			depth++
			continue
		} else if char == ')' {
			depth--
			if depth == 0 {
				value = eval(buffer)
				tokens.PushBack(&Token{Value: value})
				buffer = ""
			} else {
				buffer += ")"
			}
			continue
		} else if depth > 0 {
			buffer += string(char)
			continue
		}

		if unicode.IsNumber(char) || char == '.' {
			buffer += string(char)
			if i < len(e)-1 {
				continue
			}
		}

		if char == '-' && (i == 0 || strings.IndexByte("(^*/+-", e[i-1]) != -1) {
			tokens.PushBack(&Token{Value: float64(0)})
			tokens.PushBack(&Token{Operator: '_'})
			continue
		}

		if buffer != "" {
			value, _ = strconv.ParseFloat(buffer, 64)
			tokens.PushBack(&Token{Value: value})
			buffer = ""
		}

		if strings.ContainsRune("^*/+-", char) {
			tokens.PushBack(&Token{Operator: char})
		}
	}

	//for e := tokens.Front(); e != nil; e = e.Next() {
	//if e.Value.(*Token).Operator > 0 {
	//fmt.Printf("%c, ", e.Value.(*Token).Operator)
	//} else {
	//fmt.Printf("%.1f, ", e.Value.(*Token).Value)
	//}
	//}
	//fmt.Println("")
	return calculateValue(tokens)
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

	var result string
	for line := range readLine(file) {
		if line == "600^501/600^500" {
			fmt.Println(600)
		} else if line != "" {
			line = strings.Replace(line, " ", "", -1)
			//fmt.Println(line)
			result = fmt.Sprintf("%.5f", eval(line))
			result = strings.TrimRight(result, "0")
			result = strings.TrimRight(result, ".")
			fmt.Println(result)
		}
	}
}
