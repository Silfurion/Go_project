package main

//import "fmt"
import "math"

func dijkstra(tree [][]float64, x int) (path []int, weight []float64){
	path = make([]int, len(tree))		//Array of last nodes before goal
	weight = make([]float64, len(tree))	//Array of cost to get from first node to goal
	end := make([]bool, len(tree))	//Array of false
	compteur := 1
	previous := x
	cost := 0.0
	path[x] = x
	weight[x] = 0.0

	for i:=0;i<len(tree);i++ {
		if i != x {
			weight[i] = math.Inf(1)
		}
	}

	for compteur<len(tree) {
		min := math.Inf(1)
		location := -1
		for i:=0 ; i<len(tree) ; i++ {
			if end[i] == false {
				if cost+tree[previous][i] < weight[i] {
					path[i] = previous
					weight[i] = cost+tree[previous][i]
				}

				if cost+tree[previous][i] < min {
					min = cost+tree[previous][i]
					location = i
					end[i] = true
				}
			}
		}
		previous = location
		cost = min
		compteur++
	}
	return path, weight
}

func printPath(tree [][]float64, node int) {
	path, weight := dijkstra(tree, node)
	for i:=0;i<len(tree);i++ {
		print("Node ", i, "\nPrevious node :", path[i], "\nTotal weight of the path : ", int(weight[i]))
		print("\n\n")
	}
}

func main(){
	var tree = [][]float64{{0, 2, math.Inf(1), math.Inf(1), math.Inf(1)},	{2, 0, 3, 10, math.Inf(1)}, {math.Inf(1), 3, 0, math.Inf(1), 1}, {math.Inf(1), 10, math.Inf(1), 0, 1}, {math.Inf(1), math.Inf(1), 1, 1, 0} }
	printPath(tree, 0)
}
