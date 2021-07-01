package main

import "fmt"

func enteroProm (num1, num2, num3 int) int {
  return (num1 + num2 + num3) /3
}

func main () {
  var num1, num2, num3 int;
  fmt.Println("Introduzca primer num: ");
  fmt.Scan(&num1);
  fmt.Println("Introduzca segundo num: ");
  fmt.Scan(&num2);
  fmt.Println("Introduzca tercer num: ");
  fmt.Scan(&num3);
  promPrint := enteroProm(num1, num2, num3);
  fmt.Println("El promedio es: ", promPrint);
}
