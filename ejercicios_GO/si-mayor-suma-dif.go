package main

import "fmt"

func main () {

  var num1, num2, resultado1, resultado2 float32;

  fmt.Println("Introduzca el primer num: ");
  fmt.Scan(&num1);

  fmt.Println("Introduzca el segundo num: ");
  fmt.Scan(&num2);

  if num1 > num2 {
    resultado1 = num1 + num2;
    resultado2 = num1 - num2;
    fmt.Println("la suma es: ", resultado1);
    fmt.Println("la diferencia es: ", resultado2);
  } else {
    resultado1 = num1 * num2;
    resultado2 = num1 / num2;
    fmt.Println("El producto es: ", resultado1);
    fmt.Println("la division es: ", resultado2);
  }
  
}
