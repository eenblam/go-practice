package main

import (
    "fmt"
    "time"
)

type argRangeError struct {
    arg int
    msg string
}

func repeatInGoroutine(s string, times int) {
    //if times < 0 {
    //    return -1
    //}
    go func() {
        for i := 0; i < times; i++ {
            fmt.Println(s, ":REPEAT:", i)
            time.Sleep(1 * time.Second)
        }
    }()
}

func once(label string, message string) {
    fmt.Println(label, ":ONCE:", message)
}

func main() {
    repeatInGoroutine("A", 5)
    once("B", "First")
    once("B", "Second")

    // Allow goroutines to complete
    var input string
    fmt.Scanln(&input)
}
