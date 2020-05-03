package utils

import "fmt"

// Para que las funciones sean exportadas debe empezar con mayuscula
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

func GetIndex(a []int64, item int64) int {
	for i, v := range a {
		if v == item {
			return i
		}
	}
	return -1
}

func DeleteItemsInt(a []int64, b []int64) []int64 {
	for _, v := range b {
		index := GetIndex(a, v)
		if index != -1 {
			fmt.Println(index)
			a = Remove(a, index)
		}
	}
	return a
}

func Remove(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}
