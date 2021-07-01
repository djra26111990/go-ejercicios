package main

import "fmt"


func enteroFor (entero int) {
  if entero >= 1 {
     for x:=1; x <= entero; x++ {
       fmt.Println(x);
     }    
  } else {
    fmt.Println("Error valor negativo introducido");
  }
}


func main () {
  var entero int;

  fmt.Println("Introduzca num: ");
  fmt.Scan(&entero);
  enteroFor(entero)

}
