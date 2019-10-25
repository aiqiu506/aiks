package utils

import "math"

/**
 两浮点数进行精度带精度的比较
@params p 表示精度 如传 0.0001 表示 4位的精度相等
*/
func IsEqualFloat(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
