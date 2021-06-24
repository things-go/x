package extmath

import (
	"math"
)

// Pow is int64 type of math.Pow function.
func Pow(x, y int64) int64 {
	if y <= 0 {
		return 1
	}
	if y%2 == 0 {
		sqrt := Pow(x, y/2)
		return sqrt * sqrt
	}
	return Pow(x, y-1) * x
}

// Round 数值保留小数点,n为保留小数点位数
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

// Abs 整型取绝对值.
func Abs(number int64) int64 {
	r := number >> 63
	return (number ^ r) - r
}

// Range 根据范围创建切片,包含指定的元素.
// start: 起始元素值
// end: 末尾元素值
// 若start < end, 返回升序的数组
// 若start > end, 返回降序的数组.
func Range(start, end int) []int {
	length := int(Abs(int64(end-start))) + 1
	res := make([]int, 0, length)
	for i := 0; i < length; i++ {
		value := start
		if end > start {
			value += i
		} else {
			value -= i
		}
		res = append(res, value)
	}
	return res
}

// Range 根据范围创建数组,包含指定的元素.
// start: 起始元素值
// end: 末尾元素值
// 若start < end, 返回升序的数组
// 若start > end, 返回降序的数组.
func RangeInt64(start, end int64) []int64 {
	length := Abs(end-start) + 1
	res := make([]int64, 0, length)
	for i := int64(0); i < length; i++ {
		value := start
		if end > start {
			value += i
		} else {
			value -= i
		}
		res = append(res, value)
	}
	return res
}

func IsPowOf2(x int64) bool {
	return (x & (x - 1)) == 0
}

func NextPowOf2(x int64) int64 {
	if IsPowOf2(x) {
		return x
	}
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}
