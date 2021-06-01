package myfunc

func Defer(nums []int) int {
	var Temp = 0
	for _, v := range nums {
		Temp ^= v
	}
	return Temp
}
