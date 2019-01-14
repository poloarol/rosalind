package main

import (
    "fmt"
  )


func main(){
  wild := "GAGCCTACTAACGGGAT"
  mutant := "CATCGTAATGACGGCCT"
  count := mutation(wild, mutant)
  fmt.Println(count)
}

func mutation(s string, t string) int {

  count := 0

  if (len(s) != len(t)){
    fmt.Println("The 2 DNA strings have different length")
  }else{
    for i := 0; i < len(s); i++ {
      if(s[i] != t[i]){
        count++
      }
    }
  }
  return count
}
