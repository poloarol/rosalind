package main

import (
  "fmt"
  "bufio"
  "path/filepath"
  "strings"
  "os"
  "strconv"
)

func main(){
  var dominant, recessive, hetero, total int
  var recessive_offspring, dominant_offspring float64
  var r_r, r_h, h_h float64
  path, err := filepath.Abs("excercise/rosalind_iprb.txt")
  check(err)

  file, err := os.Open(path)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  check(err)

  population := strings.Split(parseFile(scanner), " ")
  dominant, err = strconv.Atoi(population[0])
  check(err)
  recessive, err = strconv.Atoi(population[2])
  check(err)
  hetero, err = strconv.Atoi(population[1])
  check(err)
  total = dominant + recessive + hetero

  r_r = ((float64(recessive)/float64(total)) * ((float64(recessive-1))/float64((total-1))))
  r_h = ((float64(recessive)/float64(total)) * (float64((hetero))/float64((total-1)))) * 0.5
  h_h = ((float64(hetero)/float64(total)) * (float64((hetero-1))/float64((hetero-1)))) * 0.25

  recessive_offspring = r_r + r_h + h_h

  dominant_offspring = (1 - recessive_offspring)
  fmt.Println(dominant_offspring)
}

func check(e error){
  if e != nil{
    panic(e)
  }
}

func parseFile(s *bufio.Scanner) string {
  var pop strings.Builder;
  for s.Scan(){
    pop.WriteString(s.Text());
  }
  return pop.String();
}
