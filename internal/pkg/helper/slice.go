package helper

// AppendInt64Uniq ...
func AppendInt64Uniq(slice []int64, i int64) []int64 {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
