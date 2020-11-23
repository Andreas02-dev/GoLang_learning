package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Lees argumenten van het programma uit en converteer de String input naar een integer.
	n, _ := strconv.Atoi(os.Args[1])

	nfac := faculty(n)

	fmt.Println("De faculteit van ", n, " is ", nfac)
}

// Declare faculty function:
func faculty(n int) uint64 {
	if n == 2 {
		return 2
	}

	return uint64(n) * faculty(n-1)
}

// Declare readLine function:
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	//// Voor andere operating systems dan Windows, gebruik:
	// text = strings.TrimSuffix(text, "\n")
	return text
}

// Declare scanString function:
func scanString() string {
	var s string
	fmt.Fscanf(os.Stdin, "&q", &s)
	return s
}

// Declare scanInt function:
func scanInt() int {
	var i int
	fmt.Fscanf(os.Stdin, "%d", &i)
	return i
}
