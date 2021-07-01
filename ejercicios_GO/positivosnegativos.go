package main

import "fmt"

func positivosNegativos () (pos int, neg int) {
  
  var val int;
  positivos := 0;
  negativos := 0;

  for {
    fmt.Println("Introduzca Valor");
    fmt.Scan(&val);
    if val == 0 {
      break;
    } else {
      if val > 0 {
        positivos++
        pos = positivos;
      } else {
        negativos++
        neg = negativos;
      }
    }
  }
  return
}

func main () {
  pos, neg := positivosNegativos()
  fmt.Println("El total de positivos es: ", pos);
  fmt.Println("El total de negativos es: ", neg);
}