package main

import "fmt"

func enterosVecPositivos (vec [10]int) (int, int, int) {
  
  uno := 0;
  dos := 0;
  tres := 0;

  for i := 0; i < len(vec); i++ {
    if vec[i] < 10 {
      uno++
    } else {
      if vec[i] < 100 {
        dos++
      } else {
        tres++
      }
    }
  }  
  return uno, dos, tres
}


func main () {
  vector := [10]int{5,12,2,30,1,5,85,100,200,300}
  unDig, dosDig, masDosDig := enterosVecPositivos(vector)
  fmt.Println("La cantidad de valores con 1 digito es: ", unDig);
  fmt.Println("La cantidad de valores con 2 digitos es: ", dosDig);
  fmt.Println("La cantidad de valores con mas de 2 digitos es: ", masDosDig);
}