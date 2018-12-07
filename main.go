package main

import (
    "os"
    "routely/graphs"
    "routely/reader"
    "routely/algorithm"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
  mapfile, err := os.Open("./Samples/map1.txt")
  check(err)
  queryfile, err := os.Open("./Samples/queries1.txt")
  check(err)
  intersections, roads, queries := reader.ReadData(mapfile, queryfile)
  graph := graphs.NewGraph(intersections, roads)
  circlefinder := algorithm.NewCircleFinder(intersections)
  for _, query := range queries {
    path := algorithm.BestPath(graph, query, circlefinder)
  }
}
