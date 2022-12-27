package iteration

func Repeat(item string, times int) string {
	var res string
	for i := 0; i < times; i++ {
		res += item
	}
	return res
}

func Sum(numbers ...int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	res := make([]int, len(slices))
	for i, sl := range slices {
		res[i] = Sum(sl...)
	}
	return res
}

func SumAllTails(slices ...[]int) []int {
	var res []int
	for _, sl := range slices {
		if (len(sl) == 0) {
			res = append(res, 0)
		} else {
			res = append(res, Sum(sl[1:]...))
		}
	}
	return res
}
