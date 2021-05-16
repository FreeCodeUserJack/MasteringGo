package iteration

const repeatCount = 5

func Repeat(char string, repeat int) string {

	if repeat <= 0 {
		repeat = repeatCount
	}

	var res string

	for i := 0; i < repeat; i++ {
		res += char
	}

	return res
}