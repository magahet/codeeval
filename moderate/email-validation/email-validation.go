package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
    "unicode"
)

var alphaNumericAscii = &unicode.RangeTable{
	R16: []unicode.Range16{
	    {'0', '9', 1},
	    {'A', 'Z', 1},
	    {'a', 'z', 1},
	},
}

var unQuotedSpecials = &unicode.RangeTable{
	R16: []unicode.Range16{
	    {33, 33, 1},
	    {35, 39, 1},
	    {42, 42, 1},
	    {43, 43, 1},
	    {45, 45, 1},
	    {47, 47, 1},
	    {61, 61, 1},
	    {63, 63, 1},
	    {94, 96, 1},
	    {123, 126, 1},
    },
}

var quotedSpecials = &unicode.RangeTable{
	R16: []unicode.Range16{
	    {40, 41, 1},
	    {44, 44, 1},
	    {46, 46, 1},
	    {58, 60, 1},
	    {62, 62, 1},
	    {64, 64, 1},
	    {91, 91, 1},
	    {93, 93, 1},
	},
}

var quotedSpecialsEscaped = &unicode.RangeTable{
	R16: []unicode.Range16{
	    {32, 32, 1},
	    {34, 34, 1},
	    {92, 92, 1},
	},
}

// validate domain component
func validDomain(domain string) bool {
    // too long
    if len(domain) > 253 {
        return false
    }

    domainParts := strings.Split(domain, ".")

    // not enough labels
    if len(domainParts) < 2 {
        return false
    }

    // TLD too short
    if len(domainParts[len(domainParts)-1]) < 2 {
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

// validate quoted parts of local
func validQuotedPart(part string) bool {
    for i, char := range part {
        // valid chars for quoted strings
        if unicode.Is(alphaNumericAscii, char) {
            continue
        } else if unicode.Is(unQuotedSpecials, char) {
            continue
        } else if unicode.Is(quotedSpecials, char) {
            continue
        // valid escaped chars for quoted strings
        } else if unicode.Is(quotedSpecialsEscaped, char) {
            // no escape char
            if i == 0 {
                return false
            // escape char present
            } else if part[i-1] != 0x92 {
                return false
            }
        // not a valid char
        } else {
            return false
        }
    }
    return true
}

// validate unquoted parts of local
func validUnQuotedPart(part string) bool {
    for _, char := range part {
        // not a valid char
        if unicode.Is(alphaNumericAscii, char) {
            continue
        } else if unicode.Is(unQuotedSpecials, char) {
            continue
        } else {
            return false
        }
    }
    return true
}

// validate local component
func validUser(user string) bool {
    // too long
    if len(user) > 64 {
        return false
    }

    if user[0] == '"' && user[len(user)-1] == '"' {
        return validQuotedPart(user[1:len(user)-1])
    }

    for _, part := range strings.Split(user, ".") {
        // consecutive dots or dots at start or end
        if part == "" {
            return false
        }
        if part[0] == '"' && part[len(part)-1] == '"' {
            if validQuotedPart(part[1:len(part)-1]) == false {
                return false
            }
        } else if validUnQuotedPart(part) == false {
            return false
        }
    }
    return true
}

func validateAddress(address string) bool {
    // too long
    if len(address) > 254 {
        return false
    }

    // not made up of local@domain
    index := strings.LastIndex(address, "@")
    if index <= 0 || index == len(address) - 1 {
        return false
    }

    user := address[:index]
    domain := address[index+1:]

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
        fmt.Println(validateAddress(line))
	}
}
