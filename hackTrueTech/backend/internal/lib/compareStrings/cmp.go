package compareStrings

func CmpStr(greater, lower string) int {
	if len(greater) == 0 {
		return -1
	}
	if len(lower) == 0 {
		return 1
	}

	for i := 0; ; i++ {
		if i >= len(greater)-1 {
			return -1
		}
		if i >= len(lower)-1 {
			return 1
		}

		if greater[i] > lower[i] {
			return 1
		} else if greater[i] < lower[i] {
			return -1
		}
	}
}
