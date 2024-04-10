package main

import "fmt"

type Int int

func (i Int) Double() Int {
    return i * 2
}

func main() {
    num := Int(5)
    doubled := num.Double()
    fmt.Println(doubled) // Saída: 10
}
