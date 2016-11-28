package probs

import "fmt"

type IntPoint struct {
	X, Y, Sum int
}

// set implementation for small number of items
type IntPointSet struct {
	slice []IntPoint
}

// functions

func (p1 IntPoint) Equals(p2 IntPoint) bool {
	return (p1.X == p2.X) && (p1.Y == p2.Y)
}

func (set *IntPointSet) Add(p IntPoint) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set IntPointSet) Contains(p IntPoint) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set IntPointSet) Get(p IntPoint) *IntPoint {
	for _, v := range set.slice {
		if v.Equals(p) {
			return &v
		}
	}
	return nil
}

func NewIntPointSet() IntPointSet {
	return IntPointSet{}
}

const (
	GRID_SIZE = 20
)

func main() {
	memo := NewIntPointSet()
	sum := countAllPath(0, 0, &memo)
	fmt.Println(sum)
	//fmt.Println(memo)
}

//137846528820
// 40!/20!20!
func countAllPath(i, j int, memo *IntPointSet) int {
	if i == GRID_SIZE && j == GRID_SIZE {
		return 1
	}

	var sumRight int = 0
	var sumDown int = 0
	//if there is right turn
	if j != GRID_SIZE {
		pp := memo.Get(IntPoint{i, j + 1, 0})
		if pp != nil {
			sumRight = pp.Sum
		} else {
			sumRight = countAllPath(i, j+1, memo)
			memo.Add(IntPoint{i, j + 1, sumRight})
		}
	}

	//if there is down turn
	if i != GRID_SIZE {
		ppp := memo.Get(IntPoint{i + 1, j, 0})
		if ppp != nil {
			sumDown = ppp.Sum
		} else {
			sumDown = countAllPath(i+1, j, memo)
			memo.Add(IntPoint{i + 1, j, sumDown})
		}
	}
	return sumRight + sumDown
}
