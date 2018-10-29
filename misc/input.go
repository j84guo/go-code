package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"strconv"
)

/**
 * IMO this is the correct way to read a number from stdin, read the whole line
 * first, then try to parse it, multiple tokens can be split and parsed
 * individually if needed
 */
func main() {
	fmt.Print("prompt:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	line = strings.TrimSpace(line)
	grade, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("You entered", grade)
}
