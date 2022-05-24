package utils

import "strconv"

// String2Int64	string数组转换为int数组
func String2Int64(strArr []string) []int64 {
	int64Arr := make([]int64, len(strArr))

	for idx, val := range strArr {
		int64Arr[idx], _ = strconv.ParseInt(val, 10, 64)
	}

	return int64Arr
}
