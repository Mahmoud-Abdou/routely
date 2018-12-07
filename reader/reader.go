package reader

import (
	"routely/data"
	"bufio"
	"io"
	"strconv"
)

// ReadData read test data from io stream
func ReadData(mapFile io.Reader, queryFile io.Reader) ([]*data.Intersection, []*data.Road, []*data.Query) {
	scannerd := bufio.NewScanner(mapFile)
	scannerd.Split(bufio.ScanWords)
	scannerq := bufio.NewScanner(queryFile)
	scannerq.Split(bufio.ScanWords)
	return readIntersections(scannerd), readRoads(scannerd), readQueries(scannerq)
}


func readIntersections(scanner *bufio.Scanner) []*data.Intersection {
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	ret := make([]*data.Intersection, n)
	for i := 0; i < n; i++ {
		ret[i] = &data.Intersection{}
		scanner.Scan()
		id, err := strconv.ParseUint(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		ret[i].ID = uint(id)
			scanner.Scan()
		ret[i].X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].Y, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
	}
	return ret
}

func readRoads(scanner *bufio.Scanner) []*data.Road {
	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	ret := make([]*data.Road, m)
	for i := 0; i < m; i++ {
		ret[i] = &data.Road{}
		scanner.Scan()
		from, err := strconv.ParseUint(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		ret[i].From = uint(from)
			scanner.Scan()
		to, err := strconv.ParseUint(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		ret[i].To = uint(to)
			scanner.Scan()
		ret[i].Length, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].Speed, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
	}
	return ret
}

func readQueries(scanner *bufio.Scanner) []*data.Query {
	scanner.Scan()
	q, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	ret := make([]*data.Query, q)
	for i := 0; i < q; i++ {
		ret[i] = &data.Query{}
		scanner.Scan()
		ret[i].From.X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].From.Y, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].To.X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].To.X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
			scanner.Scan()
		ret[i].WalkingRadius, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
	}
	return ret
}
