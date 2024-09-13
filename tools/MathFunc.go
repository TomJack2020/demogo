package tools

func SumNum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s

}
