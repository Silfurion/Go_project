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
var length = 0


 file, err :=os.Open("test.txt") 
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
      if count == 0 {
        leng,err := strconv.ParseInt(scanner.Text(),10,32)
        if err != nil {
          fmt.Println(err)
        }
        length = int(leng)



      }
      if count == length-1 {
          break
      }
    fmt.Println(scanner.Text())
    line :=strings.Split(scanner.Text(),";")
    fmt.Println(line)
    tab = append(tab,line)
    count++

  }
  ConvertToInt(tab,length)
    
}
func ConvertToInt(tab [][]string,length int){
  var tab2 = make([][]float64,length)
  for i:=0;i<len(tab);i++{
   for j:=0;j<len(tab);j++{
      tabb,err := strconv.ParseFloat(tab[i][j],10)
        if err!= nil {
          fmt.Println(err)
        }
      tab2[i][j]=tabb
    }
  }
  fmt.Println(tab2)


}




