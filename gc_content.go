package main

import (
  "fmt"
  "bufio"
  "path/filepath"
  "strings"
  "os"
  "sort"
)

// data structure to hold key:value pairs

type Pair struct {
  Key string
  Value float64
}

// A slice of Pairs that implement the sorting interface

type PairList []Pair
func (p PairList) Len() int {return len(p)}
func (p PairList) Less(i, j int) bool {return p[i].Value < p[j].Value}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i]}

func rankByGC(gc map[string] float64) PairList{
  pl := make(PairList, len(gc))
  i := 0
  for k, v := range gc {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

func main(){

  path, err := filepath.Abs("excercise/rosalind_gc.txt")
  check(err)

  file, err := os.Open(path)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  check(err)

  GC := parseFile(scanner)

  for i := 0; i < len(GC); i++ {
    fmt.Println(GC[i].Key, " ", GC[i].Value)
  }

}

func check(e error){
  if e != nil{
    panic(e)
  }
}

func parseFile(s *bufio.Scanner) PairList{

  var dict = make(map[string] float64)
  var key string
  var seq strings.Builder

  for s.Scan(){

    if (strings.Contains(s.Text(), ">")){
      if(len(seq.String()) != 0){
        dict[key] = calc(seq.String())
        seq.Reset()
      }
      key = strings.TrimPrefix(s.Text(), ">")
    }else{
      seq.WriteString(s.Text())
    }

  }

  return rankByGC(dict)

}

func calc(seq string) float64 {
  gc := 0

  for i := 0; i < len(seq); i++ {
    if ((string(seq[i]) == "C") || (string(seq[i]) == "G") ) {
      gc++
    }
  }

  content := (float64(gc)/float64(len(seq))) * 100

  return content
}
