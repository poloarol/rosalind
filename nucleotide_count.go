package main

import (
    "fmt"
    "strings"
  )

func main(){

  seq := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"

  a, c, g, t := 0, 0, 0, 0

  for i := 0; i < len(seq); i++ {
    if (strings.ToLower(string(seq[i])) == "a"){
      a += 1
    } else if (strings.ToLower(string(seq[i])) == "c"){
      c += 1
    }else if (strings.ToLower(string(seq[i])) == "g"){
      g += 1
    }else if (strings.ToLower(string(seq[i])) == "t"){
      t += 1
    }
  }

  output := fmt.Sprintf("A : %d, C : %d, G : %d, T : %d.", a,c,g,t)
  fmt.Println(output)
}
