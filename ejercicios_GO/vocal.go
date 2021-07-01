package main

import (
    "fmt"
    "strings"
)

func vocal (stringVal string) {
  strLowered := strings.ToLower(stringVal);
  if strLowered  == "a" || strLowered == "e" || strLowered == "i" || strLowered == "o" || strLowered == "u" {
    fmt.Println("Es una vocal.");
  } else {
    fmt.Println("No es una vocal");
  }
}

func main () {
  var stringVal string;
  fmt.Println("Introduzca una letra: ");
  fmt.Scan(&stringVal);
  vocal(stringVal)
}
