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

func main() {
	fmt.Print("Enter Test Case Number: ")

	//input
	input := bufio.NewReader(os.Stdin)
	path, _ := input.ReadString('\n')
	path = path[:len(path)-1]

	mapfile, err := os.Open("./Samples/map" + path + ".txt")
	if err != nil {
		panic(err)
	}

	queryfile, err := os.Open("./Samples/queries" + path + ".txt")
	if err != nil {
		panic(err)
	}

	//clock
	preStart := time.Now()

	intersections, roads, queries := reader.ReadData(mapfile, queryfile)
	graph := graphs.NewGraph(intersections, roads)
	circlefinder := algorithm.NewCircleFinder(intersections)

	preElapsed := time.Now().Sub(preStart)
	fmt.Println("\nPre-processing Time = ", preElapsed)

	for _, query := range queries {
		start := time.Now()
		answer := algorithm.BestPath(graph, query, circlefinder)
		elapsed := time.Now().Sub(start)

		fmt.Print("\n------------------( ͡° ͜ʖ ͡°)------------------\n\n")
		fmt.Println("Time             = ", answer.Time, "mins")
		fmt.Println("Distance         = ", answer.DrivingDistance+answer.WalkingDistance, "km")
		fmt.Println("Walking Distance = ", answer.WalkingDistance, "km")
		fmt.Println("Vehicle Distance = ", answer.DrivingDistance, "km")
		fmt.Println("Time Elapsed     = ", elapsed)
	}

	fmt.Print("\n------------------( ͡° ͜ʖ ͡°)------------------\n\n")
}
