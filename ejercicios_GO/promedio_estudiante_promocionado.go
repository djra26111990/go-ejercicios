package main

import "fmt"

func main() {

  var nota1, nota2, nota3, promedio float32;

  fmt.Print("Introduzca primera nota: ");
  fmt.Scan(&nota1);
  fmt.Print("Introduzca segunda nota: ");
  fmt.Scan(&nota2);
  fmt.Print("Introduzca primera nota: ");
  fmt.Scan(&nota3);

  promedio = (nota1 + nota2 + nota3) / 3;

  if promedio >= 7 {
    fmt.Println("Promocionado");
  } else {
    fmt.Println("No promocionado");
  }

}
