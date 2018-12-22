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

	fmt.Print("Enter testcase type (1 for normal, 2 for bonus): ")

	input := bufio.NewReader(os.Stdin)
	caseType, _ := input.ReadString('\n')
	caseType = caseType[:len(caseType)-1]

	subPath := ""

	if caseType == 2 {
		subPath = "Bonus"
	}

	fmt.Print("Enter Test Case Number: ")

	//input
	path, _ := input.ReadString('\n')
	path = path[:len(path)-1]

	mapfile, err := os.Open("./Samples" + subPath + "/map" + path + ".txt")
	if err != nil {
		panic(err)
	}

	queryfile, err := os.Open("./Samples" + subPath + "/queries" + path + ".txt")
	if err != nil {
		panic(err)
	}

	intersections, roads, queries := reader.ReadData(mapfile, queryfile, (caseType == 2))
	graph := graphs.NewGraph(intersections, roads)
	circlefinder := algorithm.NewCircleFinder(intersections)

	var totalTime time.Duration

	for _, query := range queries {
		start := time.Now()
		answer := algorithm.BestPath(graph, query, circlefinder)
		elapsed := time.Now().Sub(start)
		totalTime += elapsed

		fmt.Printf("%.2f mins\n", answer.Time)
		fmt.Printf("%.2f km\n", answer.DrivingDistance+answer.WalkingDistance)
		fmt.Printf("%.2f km\n", answer.WalkingDistance)
		fmt.Printf("%.2f km\n", answer.DrivingDistance)
		fmt.Printf("%d ms\n\n", elapsed.Round(time.Millisecond)/time.Millisecond)
	}
	fmt.Printf("%d ms\n", totalTime.Round(time.Millisecond)/time.Millisecond)
}
