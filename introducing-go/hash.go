package main

import(
	"fmt"
	"hash/crc32"
	"crypto/sha1"
	"reflect"
	"os"
	"io"
)

/*
* a hash function takes variable size a piece of data and reduces it to a fixed
* size, this is useful in data lookup (fixed width key) and detecting changes
*
* cryptographic hash functions guarantee that the hash is hard to reverse (SHA-
* 256), the rest are non-cryptographic (crc32)
*/
func crc32Demo(){
	// hasher
	h := crc32.NewIEEE()
	fmt.Println(reflect.TypeOf(h))

	// write a string converted to bytes
	h.Write([]byte("test"))

	// checksum
	v := h.Sum32()
	fmt.Println(reflect.TypeOf(v)) // uint32
	fmt.Println(v)
}

func fileHash(path string) (uint32, error){
	file, err := os.Open(path)

	if err != nil{
		return 0, err
	}else{
		defer file.Close()
	}

	// write file into the hasher
	h := crc32.NewIEEE()
	_, err = io.Copy(h, file)

	if err != nil{
		return 0, err
	}

	return h.Sum32(), nil
}

func sha1Demo(){
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

func main(){
	crc32Demo()
	fmt.Println(fileHash("maps.go"))

	sha1Demo()
}
