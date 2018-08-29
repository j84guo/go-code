package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
)

type HomePageSize struct{
    URL string
    Size int
}

/*
 * question: if a goroutine panics, who recovers, the caller?
 */
func main(){
    urls := []string{
        "http://www.apple.com",
        "http://www.amazon.com",
        "http://www.google.com",
        "http://www.microsoft.com",
    }

    results := make(chan HomePageSize)

    for _, url := range urls{
        go func(url string){
            res, err := http.Get(url)
            if err != nil{
                panic(err)
            }
            defer res.Body.Close()

            bs, err := ioutil.ReadAll(res.Body)
            if err != nil{
                panic(err)
            }

            results <-HomePageSize{
                URL: url,
                Size: len(bs),
            }
        }(url)
    }

    var biggest HomePageSize

    for range urls{
        res := <-results
        if res.Size > biggest.Size{
            biggest = res
        }
    }

    fmt.Println("biggest home page:", biggest.URL, biggest)
}
