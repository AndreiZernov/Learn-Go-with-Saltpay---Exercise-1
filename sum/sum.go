package sum

type Numbers []int

func (n Numbers) add() int {
	sum := 0
	for _, number := range n {
		sum += number
	}
	return sum
}
