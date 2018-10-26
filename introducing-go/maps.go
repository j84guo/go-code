package main

import "fmt"

// maps are unordered collections of key value pairs (abstract data type)
// generally implemented as hash tables (ordered maps may use red black trees)
// also known as associative array or dictionary
func main() {
	fmt.Println("starting demo...")
	mapDemo()
	deleteDemo()
	zeroDemo()
	okDemo()
	initDemo()
	nestedDemo()
}

func initDemo() {
	// like with arrays maps can be inferred and initialized at the same time
	x := map[string]string{
		"kit": "kat",
		"tob": "lerone",
	}

	x["cara"] = "mel"
	fmt.Println(x)
}

func nestedDemo() {
	e := map[string]map[string]string{
		"H": map[string]string{
			"name":  "hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "helium",
			"state": "gas",
		},
	}

	fmt.Println(e)
}

func mapDemo() {
	// map[<key type>]<value type>
	// the make function needs to be used to initialize a map
	// maps need to be initialized before they are used
	var x map[string]int = make(map[string]int)
	x["one"] = 1

	// note the _ placeholder and the use of multiple statements in an if
	// conditional, the last of which is evaluated
	if _, ok := x["one"]; ok {
		fmt.Println(x)
	} else {
		fmt.Println("not ok")
	}
}

func deleteDemo() {
	x := make(map[string]string)
	x["one"] = "un"
	x["two"] = "deux"

	// the built in function delete removes an entry from a map
	delete(x, "one")

	fmt.Println(x)
}

func zeroDemo() {
	// maps return the zero value for that type if the key cannot be found
	// for strings, this is an empty string
	x := make(map[string]string)
	fmt.Println(x["test"])
}

func okDemo() {
	// accessing a map can return two values, the second of which indicates if it
	// succeeded or not
	x := make(map[string]string)
	v, ok := x["test"]
	fmt.Println(v, ok)
}
