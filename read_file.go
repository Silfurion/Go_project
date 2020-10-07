package main 

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"


)

func main() {
var count = 0
var tab [][]string


 file, err :=os.Open("test.txt") 
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
      if count == len(tab)-1 {
          break
      }
    fmt.Println(scanner.Text())
    line :=strings.Split(scanner.Text(),";")
    fmt.Println(line)
    tab = append(tab,line)
    count++

  }
  ConvertToInt(tab)
    
}
func ConvertToInt(tab [][]string){
  var tab2 [][]int
  var tabb []int
    for i :=0;i<len(tab);i++{
      for j:=0;j<len(tab);j++{
        nb,err := strconv.Atoi(tab[i][j])
        if err != nil{
          fmt.Println(err)
        }
        tabb=append(tabb,nb)
      }
    tab2=append(tab2,tabb)
    tabb = tabb[:1]
    } 
  fmt.Println(tab2)




}




