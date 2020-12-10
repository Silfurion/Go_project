package main

//import "fmt"
import (
	"math"
	"./read_file"
	"./WriteFile"
)

func dijkstra(tree [][]float64, x int) (sample[][]float64) {
	path := make([]int, len(tree))		//Array of last nodes before goal
	weight := make([]float64, len(tree))	//Array of cost to get from first node to goal
	end := make([]bool, len(tree))	//Array of false to know if the node has already been checked
	compteur := 0
	previous := x
	cost := 0.0
	path[x] = x
	weight[x] = 0.0

	for i:=0;i<len(tree);i++ { //Initialization
		if i != x {
			weight[i] = math.Inf(1)
		}
	}

	for compteur<len(tree) {
		min := math.Inf(1)
		location := -1
		for i:=0 ; i<len(tree) ; i++ {
			if end[i] == false { //If the node isn't locked yet, check the weight of the shortest path
				if cost+tree[previous][i] < weight[i] { //Saving the state of each shortest path per node.
					path[i] = previous
					weight[i] = cost+tree[previous][i]
				}

				if weight[i] < min { //Acknoweldging the shortest path
					min = weight[i]
					location = i
				}
			}
		}
		end[location] = true
		previous = location
		cost = min
		compteur++
	}
	sample = make([][]float64, len(tree)+1)
	for i:= range sample{
		sample[i] = make([]float64, 2)
	}
	sample[len(sample)-1][0] = float64(x)

	for i:=0;i<len(sample)-1;i++{
		sample[i][0] = weight[i]
		sample[i][1] = float64(path[i])
	}
	return sample
}

func loadDijkstra(tree [][]float64, node int, done chan string, data chan [][]float64) {
	sample := dijkstra(tree, node)
	data <- sample
	done <- "done"
}

func loadData(tree [][]float64,data chan [][]float64) (result [][][]float64){
	result = make([][][]float64, len(tree))
	for i:= range result{
		result[i] = make([][]float64, len(tree))
		for j:= range result{
			result[i][j] = make([]float64, 2)
		}
	}
	
	loop := true
	for loop{
		select{
		case slice := <- data:
			for i:=0;i<len(result);i++{
				result[int(slice[len(slice)-1][0])][i][0] = slice[i][0] //Distance
				result[int(slice[len(slice)-1][0])][i][1] = slice[i][1] //Previous node
			}
		default:
			loop = false
		}
	}
	return result
}

func printTree(tree [][][]float64){
	for i:=0;i<len(tree);i++{
		print("Node ", i, " :\n")
		for j:=0;j<len(tree);j++{
			print("to ", j, " : distance=", int(tree[i][j][0]), " and previous node is ", int(tree[i][j][1]), ".\n")
		}
		print("\n\n")
	}
}


func main(){
	//var tree = [][]float64{{0, 2, math.Inf(1), math.Inf(1), math.Inf(1)},	{2, 0, 3, 10, math.Inf(1)}, {math.Inf(1), 3, 0, math.Inf(1), 1}, {math.Inf(1), 10, math.Inf(1), 0, 1}, {math.Inf(1), math.Inf(1), 1, 1, 0} }
	var tree = read_file.ReadFile()
	
	tokens := make(chan int, len(tree))
	done := make(chan string, len(tree))
	data := make(chan [][]float64, len(tree))
	
	for i:=0;i<len(tree);i++ {
		tokens <- i
	}

	count := len(tree)
	routines := 0
	for count > 0 {
		if routines < 4 {
			select{
				case token := <- tokens :
					go loadDijkstra(tree, token, done, data)
					routines++
				default:
					//do nothing
			}
		}
		if <- done == "done" {
			routines--
			count--
		}
	}
	result := loadData(tree, data)
	WriteFile.WriteFile(result)
	printTree(result)
}
