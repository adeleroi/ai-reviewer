package experiments

func DotProdV1(v1, v2 []int) int {
	var res = 0
	for idx := range v1 {
		res += v1[idx] * v2[idx]
	}
	return res
}

func DotProdV2(v1, v2 []int) int {
	if len(v1) != len(v2) {
		return 0
	}
	return DotProdV1(v1, v2)
}
