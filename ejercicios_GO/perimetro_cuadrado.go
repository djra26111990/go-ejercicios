package main

import "fmt"

func perimetro (lado float32) float32 {
  return lado + lado + lado + lado
}

func main () {
  var lado float32;
  fmt.Print("Introduzca valor de un lado de un cuadrado: ");
  fmt.Scan(&lado);
  perimetroResult := perimetro(lado);
  fmt.Println("El perimetro del cuadrado es: ", perimetroResult);
}
