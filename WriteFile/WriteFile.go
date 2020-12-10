package WriteFile

import(
  "fmt"
  "bufio"
  "os"
  "strconv"
)
func ConvertToString(tab [][]float64) string {
	var intertab []string
	for i := 0 ; i<len(tab) ; i++ {
		intertabf :=  strconv.FormatFloat(tab[i][0],'G',-1,64)+";"+strconv.FormatFloat(tab[i][1],'G',-1,64)
		intertab = append(intertab,intertabf)

	}
	fmt.Println(intertab)
	return ConvertTabStringToString(intertab)
}

func ConvertTabStringToString (tabs []string) string {
	var str string 
	for i:= 0 ; i<len(tabs)-1 ;i++ {
		str+=tabs[i]+" "
	}
	str+=tabs[len(tabs)-1]
	str+="\n"
	fmt.Println(str)
	return str



}

func WriteFile (tab [][][]float64){
	file,err := os.Create("./End.txt")
	if err != nil {
		fmt.Println("File reading error", err) 
	}
	writer := bufio.NewWriter(file)
	tabstr,err := writer.WriteString(strconv.FormatInt(int64(len(tab)),16)+"\n")
	fmt.Println(tabstr)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0 ; i<len(tab) ; i++ {
		tabstr,err :=writer.WriteString(ConvertToString(tab[i]))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tabstr)
		err2 :=writer.Flush()
		if err2 != nil {
			fmt.Println(err2)
		}
	}

}

