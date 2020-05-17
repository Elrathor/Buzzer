package util

func GetIndexOfString(value string, slice *[]string) (index int64) {
	for i, element := range *slice {
		index = int64(i)
		if element == value {
			return
		}
	}

	//-1 to indicate that no element has been found
	index = -1
	return
}
