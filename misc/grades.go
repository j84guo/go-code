package main

import (
	"fmt"
	"os"
	"io"
)

/**
 * Computes a letter score from grade percentage
 *
 * Note: Scanf leaves unread data in the user buffer after an error, so 
 * subsequent reads fail, also there's strange behaviour with EOF getting
 * cleared automatically
 */
func main() {
	var g int

	for {
		_, e := fmt.Scanf("%d", &g)

		if e != nil {
			if e == io.EOF {
				break
			}

			fmt.Fprintf(os.Stderr, e.Error() + "\n")
			continue
		} else if g < 0 || g > 100 {
			fmt.Fprintf(os.Stderr, "invalid grade\n")
			continue
		}

		fmt.Println(letterGrade(g))
	}
}

func letterGrade(g int) string {
	if g < 60 {
		return "F"
	} else if g < 70 {
		return "C"
	} else if g < 80 {
		return "B"
	}

	return "A"
}
