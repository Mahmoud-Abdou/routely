package main

import (
	"bufio"
	"fmt"
	"os"
	"routely/algorithm"
	"routely/graphs"
	"routely/reader"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//input
	input := bufio.NewReader(os.Stdin)
	path, _ := input.ReadString('\n')
	path = path[:len(path)-1]

	//clock
	preStart := time.Now()

	mapfile, err := os.Open("./Samples/map" + path + ".txt")
	check(err)

	queryfile, err := os.Open("./Samples/queries" + path + ".txt")
	check(err)

	intersections, roads, queries := reader.ReadData(mapfile, queryfile)
	graph := graphs.NewGraph(intersections, roads)
	circlefinder := algorithm.NewCircleFinder(intersections)

	preElapsed := time.Now().Sub(preStart)
	fmt.Println("Pre-processing Time = ", preElapsed)

	for _, query := range queries {
		start := time.Now()
		answer := algorithm.BestPath(graph, query, circlefinder)
		elapsed := time.Now().Sub(start)

		fmt.Println()
		fmt.Println("-----------------------( ͡° ͜ʖ ͡°)-----------------------")
		fmt.Println()
		fmt.Println("Time =  ", answer.Time, " mins")
		fmt.Println("Distance = ", answer.DrivingDistance+answer.WalkingDistance, "km")
		fmt.Println("Walking Distance = ", answer.WalkingDistance, "km")
		fmt.Println("Vehicle Distance = ", answer.DrivingDistance, "km")
		fmt.Println("Elapse Time = ", elapsed)
		fmt.Println("------------------------------------------------------")
	}

}
