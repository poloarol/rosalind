package main

import (
  "fmt"
  "strings"
)

func main(){


  seq := "ACTGCTA"

  const A string = "A"
  const G string = "G"
  const C string = "C"
  const T string = "T"

  var str strings.Builder

  for i := len(seq)-1; i > 0; i-- {

    if (strings.ToLower(string(seq[i])) == "a"){
      str.WriteString(T)
    } else if (strings.ToLower(string(seq[i])) == "c"){
      str.WriteString(G)
    }else if (strings.ToLower(string(seq[i])) == "g"){
      str.WriteString(C)
    }else if (strings.ToLower(string(seq[i])) == "t"){
      str.WriteString(A)
    }

  }

  fmt.Println(str.String())
}
