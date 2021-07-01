package main

import "fmt"

func main() {
    var num int
    fmt.Println("Ingrese un valor entero comprendido entre 1 y 99: ")
    fmt.Scan(&num)
    if num < 10 {
        fmt.Println("Tiene un dígito")
    } else {
        fmt.Println("Tiene dos dígitos")
    }
}
