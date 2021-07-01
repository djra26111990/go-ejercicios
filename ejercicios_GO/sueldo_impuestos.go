package main

import "fmt"

func main () {

  var sueldo float32;
  fmt.Print("Introduzca sueldo: ");
  fmt.Scan(&sueldo);
  if sueldo > 3000 {
    fmt.Print("Esta persona debe abonar impuestos");
  } else {
    fmt.Print("Esta persona no debe abonar impuestos");
  }

}
