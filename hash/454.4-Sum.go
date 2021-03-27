package main

// 四数相加 II
func fourSumCount(A, B, C, D []int) int {
	var (
		// A+B之和为键，出现的次数为值
		sumAB  = make(map[int]int)
		resutl int
	)
	// 双重循环计算A+B的和
	for _, a := range A {
		for _, b := range B {
			sumAB[a+b]++
		}
	}
	// 计算C+D-(A+B)=0的数量
	for _, c := range C {
		for _, d := range D {
			if k, ok := sumAB[0-(c+d)]; ok {
				resutl += k
			}
		}
	}
	return resutl
}
