package reader

import(
	"fmt"
	"strings"
  "testing"
)

func TestSum(t *testing.T) {
  const map1 = "6 0 0.25 0.33 1 0.73 0.47 2 1.00 0.54 3 0.34 0.60 4 0.81 0.20 5 1.07 0.28 7 0 1 0.50 20 0 3 0.28 80 3 4 0.62 80 1 4 0.28 40 1 2 0.28 40 4 5 0.27 60 2 5 0.27 60"
  const q1 = "1 0.16 0.21 1.09 0.44 300"
  intr, rood, queries := ReadData(strings.NewReader(map1), strings.NewReader(q1))
  if len(intr) != 6 {
    t.Errorf("error size of intersections")
  }
  if len(rood) != 7 {
    t.Errorf("error size of roods")
  }
  if len(queries) != 1 {
    t.Errorf("error size of queries")
  }
  fmt.Println(len(intr))
  for i := 0; i < len(intr); i++ {
    fmt.Println(intr[i])
  }
  fmt.Println(len(rood))
  for i := 0; i < len(rood); i++ {
    fmt.Println(rood[i])
  }
  fmt.Println(len(queries))
  for i := 0; i < len(queries); i++ {
    fmt.Println(queries[i])
  }
}
