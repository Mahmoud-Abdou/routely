package data

// Point represents a 2D point
type Point struct {
	X float64
	Y float64
}

// Intersection holds a vertex's data
type Intersection struct {
	ID int
	Point
}

// Road holds road data
type Road struct {
	From   int
	To     int
	Length float64
	Speed  float64
	Weight float64
}

// DijkstraNode holds dijkstra data in heap
type DijkstraNode struct {
	To     int
	Weight float64
	Index  int
}

// CircleVertex holds vertex data in respect of circle
type CircleVertex struct {
	*Intersection
	Distance float64
}

// Query holds a query's data
type Query struct {
	From          Point
	To            Point
	WalkingRadius float64
}

// Path holds path distance and points
type Path struct {
	Time            float64
	WalkingDistance float64
	DrivingDistance float64
	Steps           []*Intersection
}
