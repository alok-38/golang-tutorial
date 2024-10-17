package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var min, max int

    fmt.Print("Enter the minimum number: ")
    fmt.Scan(&min)
    fmt.Print("Enter the maximum number: ")
    fmt.Scan(&max)

    if min >= max {
        fmt.Println("Invalid range. The minimum must be less than the maximum.")
        return
    }

    randomNum := rand.Intn(max-min) + min
    fmt.Println("My favorite number is", randomNum)
}

