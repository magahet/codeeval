package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
    "unicode"
)

// validate domain component
func validDomain(domain string) bool {
    domainParts := strings.Split(domain, ".")

    // not enough labels
    if len(domainParts) < 2 {
        return false
    }

    for _, label := range domainParts {
        // label contains invalid characters
        for _, char := range label {
            if char == '-' {
                return true
            }
            if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
                return false
            }
        }

        // label starts or ends in hyphen
        if label[0] == '-' || label[len(label)-1] == '-' {
            return false
        }
    }

    return true
}

// validate local component
func validUser(domain string) bool {
    return true
}

func processLine(address string) bool {
    // too long
    if len(address) > 254 {
        return false
    }

    // not made up of local@domain
    parts := strings.Split(address, "@")
    if len(parts) != 2 {
        return false
    }

    user := parts[0]
    domain := parts[1]

    // local or domain too long
    if len(user) > 64 || len(domain) > 253 {
        return false
    }

    if validDomain(domain) == false {
        return false
    }

    if validUser(user) == false {
        return false
    }

    return true
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
        if len(line) == 0 {
            continue
        }
        //fmt.Println(line)
        fmt.Printf("%v, %s\n", processLine(line), line)
        //fmt.Println("")
	}
}
