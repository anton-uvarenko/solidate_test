package math

func SumOf(d int) int {
	sum := 0
	for d > 0 {
		sum += d % 10
		d = (d - d%10) / 10
	}
	return sum
}
