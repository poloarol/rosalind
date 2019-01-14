package main

import (
  "fmt"
  "strings"
)

func main(){

  seq := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"

  const A string = "A"
  const G string = "G"
  const C string = "C"
  const U string = "U"

  var rna strings.Builder

  for i := 0; i < len(seq); i++ {
    if (strings.ToLower(string(seq[i])) == "a"){
      rna.WriteString(A)
    } else if (strings.ToLower(string(seq[i])) == "c"){
      rna.WriteString(C)
    }else if (strings.ToLower(string(seq[i])) == "g"){
      rna.WriteString(G)
    }else if (strings.ToLower(string(seq[i])) == "t"){
      rna.WriteString(U)
    }
  }

  fmt.Println(rna.String())

}
