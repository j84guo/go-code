package main

import(
  "fmt"
  "os"
  "io/ioutil"
)

/*
before looking at files, we must understand Go's io package
the io package contains some functions and many interfaces used in other
packages

e.g. Reader and Writer, Readers read via the Read() method and writers
write via the Write() method, many functions in Go take readers and writers,
like the copy function func Copy(dst Writer, src Reader) (written int64, err error)

to read or write to a []byte or string, use the Buffer struct from the bytes
package

e.g. var buf bytes.Buffer
buf.Write([]bytes("test"))

Buffers don't have to be initialized and support Reader and Writer interfaces
you can convert into a slice of bytes by calling the buf.Bytes() method
note strings.NewReader reads from a string efficiently
*/
func main(){
  openStatReadDemo()
  readFileDemo()
  createDemo()
}

func openStatReadDemo(){
  // open a file using the open function of the os package
  file, err := os.Open("hello.go")
  if err != nil{
    fmt.Println("open: could not open file"); return
  }else{
    // Golang's approach to RAII
    defer file.Close()
  }

  // file info
  stat, err := file.Stat()
  if err != nil{
    fmt.Println("Stat: error getting file info"); return
  }

  // read into buffer
  buf := make([]byte, stat.Size())
  _, err = file.Read(buf)
  if(err != nil){
    fmt.Println("Read: could not read file"); return
  }

  // binary to string
  str := string(buf)
  fmt.Println(str)

  // file.Close() is called here
}

func readFileDemo(){
  // open read and then close file
  bs, err := ioutil.ReadFile("maps.go")
  if err != nil{
    fmt.Println("ReadFile: could not read file"); return
  }

  // binary to string
  str := string(bs)
  fmt.Println(str)
}

// creates a file and returns a file object
// todo: default permission bits? which file mode? (r/w/a, seems to be r+/w+)
func createDemo(){
  file, err := os.Create("create.txt")
  if err != nil{
    fmt.Println("Create: could not create file"); return
  }else{
    // ensure close
    defer file.Close()
  }

  file.WriteString("create this file using Create() function from os package\n")
}
