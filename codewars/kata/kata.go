package kata

func Between(a, b int) []int {
	if a > b {
		return nil
	}

	var resArr []int

	for i := a; i <= b; i++ {
		resArr = append(resArr, i)
	}

	return resArr

}
