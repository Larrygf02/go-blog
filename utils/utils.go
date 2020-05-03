package utils

import "fmt"

func AppendInt(a []int64, b []int64) []int64 {
	check := make(map[int64]int)
	d := append(a, b...)
	res := make([]int64, 0)
	for _, val := range d {
		check[val] = 1
	}
	for n, _ := range check {
		res = append(res, n)
	}
	fmt.Println(res)
	return res
}
