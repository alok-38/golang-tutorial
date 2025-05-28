package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println("Go Quote:", quote.Go())
    fmt.Println("Hello Quote:", quote.Hello())
    fmt.Println("Optional Quote:", quote.Opt())
    fmt.Println("Glass Quote:", quote.Glass())
}

