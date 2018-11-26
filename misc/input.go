package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
)

/**
 * IMO this is the correct way to read a number from stdin, read the whole line
 * first, then try to parse it, multiple tokens can be split and parsed
 * individually if needed
 */
func main() {
	fmt.Print("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	line, e := reader.ReadString('\n')
	if e != nil {
		if e == io.EOF {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, e.Error())
		os.Exit(1)
	}

	f, e := strconv.ParseFloat(line[0:len(line)-1], 64)
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}
	fmt.Println("You entered", f)
}
