package data

// Point represents a 2D point
type Point struct {
	X float64
	Y float64
}

// Intersection holds a vertex's data
type Intersection struct {
	ID uint
	Point
}

// Road holds road data
type Road struct {
	From   uint
	To     uint
	Length float64
	Speed  float64
	Weight float64
}

// Query holds a query's data
type Query struct {
	From          Point
	To            Point
	WalkingRadius float64
}

// Path holds path distance and points
type Path struct {
	Length float64
	Steps  []*Intersection
}
