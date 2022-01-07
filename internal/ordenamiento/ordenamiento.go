package ordenamiento

import "sort"

func orderSlice(a []int) []int {
	sort.Ints(a)
	return a
}
