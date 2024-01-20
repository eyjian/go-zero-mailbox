// Package logic
// Written by yijian on 2024/01/20
package logic

import "strconv"

func getInStr(lettersIdList []int64) string {
	inStr := ""
	for index, letterId := range lettersIdList {
		if index == 0 {
			inStr = strconv.FormatInt(letterId, 10)
		} else {
			inStr = inStr + "," + strconv.FormatInt(letterId, 10)
		}
	}
	return inStr
}
