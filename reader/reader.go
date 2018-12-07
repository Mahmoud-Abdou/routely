package reader

import (
	"routely/data"
	"bufio"
	"io"
)

// ReadData read test data from io stream
func ReadData(file io.Reader) ([]*data.Intersection, []*data.Road, []*data.Query) {
	return nil, nil, nil
}

func readIntersections(scanner *bufio.Scanner) []*data.Intersection {
	return nil
}

func readRoads(scanner *bufio.Scanner) []*data.Road {
	return nil
}

func readQueries(scanner *bufio.Scanner) []*data.Query {
	return nil
}
