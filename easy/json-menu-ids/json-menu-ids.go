package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func processLine(line string) int {
	var v map[string]interface{}
	json.Unmarshal([]byte(line), &v)
	sum := 0

	if menuIntr, ok := v["menu"]; ok {
		if menuIntr == nil {
			return sum
		}
		menu := menuIntr.(map[string]interface{})
		if itemsIntr, ok := menu["items"]; ok {
			if itemsIntr == nil {
				return sum
			}
			items := itemsIntr.([]interface{})
			for _, itemIntr := range items {
				if itemIntr == nil {
					continue
				}
				item := itemIntr.(map[string]interface{})
				if _, ok := item["label"]; ok {
					if idIntr, ok := item["id"]; ok {
						if idIntr != nil {
							sum += int(idIntr.(float64))
						}
					}
				}
			}
		}
	}
	return sum
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
