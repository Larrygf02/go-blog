package utils

func appendInt(a []int, b []int) []int {
	check := make(map[int]int)
	d := append(a, b...)
	res := make([]int, 0)
	for _, val := range d {
		check[val] = 1
	}
	for n, _ := range check {
		res = append(res, n)
	}
	return res
}
