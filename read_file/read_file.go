package read_file

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"


)

func ReadFile() [][]float64 {


var tab [][]float64
var count = 0
var length = 0


 file, err :=os.Open("test.txt") 
    if err != nil {
        fmt.Println("File reading error", err)
        return nil
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
        count++
      }else {
        break
      }

     }
     fmt.Println(tab)
     return tab
}
func ConvertToInt(tab []string) []float64{
  var intertab []float64
  for  i := 0 ; i<len(tab) ; i++ {
    intertabf,err := strconv.ParseFloat(tab[i],64)
    if err != nil {
    	if tab[i] == "inf" {
    		intertabf = math.Inf(1)
    	
    	}else { 	
        	fmt.Println(err)
        }
    }
    intertab = append(intertab,intertabf)
  }
  return intertab
  
}




