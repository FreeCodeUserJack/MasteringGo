package main

func Sum(nums []int) (sum int) {
	// for i := 0; i < 5; i++ {
	// 	sum += nums[i]
	// }

	for _, number := range nums {
		sum += number
	}
	return
}

func SumAll(inputs ...[]int) (sums []int) {

	// lengthOfInputs := len(inputs)
	// result := make([]int, lengthOfInputs)

	for _, slice := range inputs {
		sums = append(sums, Sum(slice))
	}

	return
}

func SumAllTails(input ...[]int) (res []int) {
	for _, slice := range input {
		if len(slice) > 0 {
			res = append(res, Sum(slice[1:]))
		} else {
			res = append(res, 0)
		}
	}
	return
}