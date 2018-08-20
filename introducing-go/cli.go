package main

import(
	"fmt"
	"flag"
	"time"
	"math/rand"
)

func main(){
	max := flag.Int("max", 6, "maximum value")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(*max))

	// non flag arguments and unregistered flags
	fmt.Println(flag.Args())
}
