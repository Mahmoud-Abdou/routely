package main

import (
    "os"
    "routely/graphs"
    "routely/reader"
    "routely/algorithm"
    "fmt"
    "bufio"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
  //input
  input := bufio.NewReader(os.Stdin)
  mapPath, _ := input.ReadString('\n')
  queryPath := mapPath
  mapPath = mapPath[:len(mapPath)-1]
  queryPath = queryPath[:len(queryPath)-1]

  mapfile, err := os.Open("./Samples/map" + mapPath + ".txt")
  check(err)
  queryfile, err := os.Open("./Samples/queries" + queryPath + ".txt")
  check(err)
  intersections, roads, queries := reader.ReadData(mapfile, queryfile)
  graph := graphs.NewGraph(intersections, roads)
  circlefinder := algorithm.NewCircleFinder(intersections)
  for _, query := range queries {
    answer := algorithm.BestPath(graph, query, circlefinder)
    fmt.Println("----------------------( ͡° ͜ʖ͡°)--------------------------")
    fmt.Println()
    fmt.Println("Time =  ", answer.Time, " mins")
    fmt.Println("Distance = ", answer.DrivingDistance + answer.WalkingDistance, "km")
    fmt.Println("Walking Distance = ", answer.WalkingDistance, "km")
    fmt.Println("Vehicle Distance = ", answer.DrivingDistance, "km")
    fmt.Println("--------------------------------------------------------")
  }

}
