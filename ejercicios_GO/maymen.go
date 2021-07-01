package main

import "fmt"

func mayorMenor(v1, v2 int) (may int, men int) {
    if v1 > v2 {
        may = v1
        men = v2
    } else {
        may = v2
        men = v1
    }
    return
}

func main() {
    var x1, x2 int
    fmt.Print("Ingrese primer valor:")
    fmt.Scan(&x1)
    fmt.Print("Ingrese segundo valor:")
    fmt.Scan(&x2)
    may, men := mayorMenor(x1, x2)
    fmt.Println("El mayor es:", may)
    fmt.Println("El menor es:", men)
}