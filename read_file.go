package main 

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"


)

func main() {


var tab [][]string
var count = 0
var length = 0


 file, err :=os.Open("test.txt") 
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
      if count == 0 {
        count++
        leng,err := strconv.ParseInt(scanner.Text(),10,32)
        if err != nil {
          fmt.Println(err)
        }
        length = int(leng)



      
      } else if count == length+1 {
          break
      
      } else{
        fmt.Println(scanner.Text())
        line :=strings.Split(scanner.Text(),";")
        fmt.Println(line,len(line))
        tab = append(tab,line)
        count++
      }

  }
  fmt.Println(tab,len(tab))
  ConvertToInt(tab,length)
    
}
func ConvertToInt(tab [][]string,length int){
  var tab2 [][]float64
  var intertab []float64
  for i:=0;i<len(tab);i++{
   for j:=0;j<len(tab);j++{
      tabb,err := strconv.ParseFloat(tab[i][j],10)
        if err!= nil {
          fmt.Println(err)
        }
        intertab = append(intertab,tabb)
    }
      tab2 = append(tab2,intertab)
      intertab = nil
    }
    fmt.Println(tab2)
}








