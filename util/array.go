package util

func Contains(arr []string, searchStr string) bool {
	for _, str := range arr {
		if str == searchStr {
			return true
		}
	}

	return false
}
