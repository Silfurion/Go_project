package main 

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"


)

func main() {


var tab [][]int
var count = 0
var length = 0


 file, err :=os.Open("test.txt") 
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    reader := bufio.NewReader(file)
     
    for {
      if count == 0 {
        read, isPrefix, err:= reader.ReadLine()
        if isPrefix {
          fmt.Println("error")
          break
        }
        if err != nil{
          fmt.Println(err)
          break
        }
        length,err = strconv.Atoi(string(read))
        if err != nil {
          fmt.Println(err)
          break
        }
        count++
      }
      if count < int(length)+1 {
        read,isPrefix,err := reader.ReadLine()
        if isPrefix  {
          break
        }
        if err != nil {
          fmt.Println(err)
          break
        }
        intertab := strings.Split(string(read)," ")
        intertabf := ConvertToInt(intertab)
        tab = append(tab,intertabf)
        if err != nil {
          fmt.Println(err)
          break
        }
       // tab = append(tab,intertab)
        count++
      }else {
        break
      }

     }
     fmt.Println(tab)
}
func ConvertToInt(tab []string) []int{
  var intertab []int
  for  i := 0 ; i<len(tab) ; i++ {
    intertabf,err := strconv.Atoi(tab[i])
    if err != nil {
      fmt.Println(err)
    }
    intertab = append(intertab,intertabf)
  }
  return intertab
  
}








