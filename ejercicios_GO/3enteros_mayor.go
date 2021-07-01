package main

import "fmt"

func mayorNum (num1, num2, num3 int) {
  if num1 > num2 {
        if num1 > num3 {
            fmt.Print("El mayor es: ", num1)
        } else {
            fmt.Print("El mayor es: ", num3)
        }        
    } else {
        if num2 > num3 {
            fmt.Print("El mayor es: ", num2)
        } else {
            fmt.Print("El mayor es: ", num3)
        }
    }
}

func main () {
  var num1, num2, num3 int;
  fmt.Println("Introduzca el primer num: ");
  fmt.Scan(&num1);
  fmt.Println("Introduzca el segundo num: ");
  fmt.Scan(&num2);
  fmt.Println("Introduzca el tercer num: ");
  fmt.Scan(&num3);

  mayorNum(num1, num2, num3)
} 
