package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello" + " " + "world")

	fmt.Println(
		containsDemo("rowboat", "boat"))

	fmt.Println(
		countDemo("abababab", "a"))

	fmt.Println(
		hasPrefixDemo("waterloo", "water"))

	fmt.Println(
		hasSuffixDemo("circuits", "ts"))

	fmt.Println(
		indexDemo("waterloo", "loo"))

	fmt.Println(
		joinDemo([]string{"testing", "one", "two", "three"}, "."))

	fmt.Println(
		repeatDemo("a", 3))

	fmt.Println(replaceDemo(
		"jackson", "a", "TEST", -1))

	fmt.Println(
		splitDemo("a,b,c", ","))

	fmt.Println(
		toLowerDemo("JAC"))

	fmt.Println(
		toUpperDemo("jac"))

	binaryDemo()
}

// search for a smaller string in a bigger string
func containsDemo(a string, b string) bool {
	return strings.Contains(a, b)
}

// count the occurences of a smaller string in a bigger string
func countDemo(a string, b string) int {
	return strings.Count(a, b)
}

// check for a prefix
func hasPrefixDemo(a string, b string) bool {
	return strings.HasPrefix(a, b)
}

// check for a suffix
func hasSuffixDemo(a string, b string) bool {
	return strings.HasSuffix(a, b)
}

// find index
func indexDemo(a string, b string) int {
	return strings.Index(a, b)
}

// join a list of strings
func joinDemo(l []string, s string) string {
	return strings.Join(l, s)
}

// repeat a string
func repeatDemo(s string, i int) string {
	return strings.Repeat(s, i)
}

// replace a substring some number of times, -1 for as many times as possible
func replaceDemo(src string, sub string, str string, i int) string {
	return strings.Replace(src, sub, str, i)
}

// split a string
func splitDemo(src string, sep string) []string {
	return strings.Split(src, sep)
}

// to lower case
func toLowerDemo(s string) string {
	return strings.ToLower(s)
}

// to upper case
func toUpperDemo(s string) string {
	return strings.ToUpper(s)
}

// to slice of bytes (binary data)
func binaryDemo() {
	bin := []byte("jackson")
	fmt.Println(bin)

	str := string(bin)
	fmt.Println(str)
}
