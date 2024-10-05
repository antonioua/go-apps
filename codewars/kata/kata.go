package kata

func HoopCount(n int) string {
	res := ""
	if n >= 10 {
		res = "Great, now move on to tricks"
	} else {
		res = "Keep at it until you get it"
	}
	return res
}
