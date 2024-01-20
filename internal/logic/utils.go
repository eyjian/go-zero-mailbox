// Package logic
// Written by yijian on 2024/01/20
package logic

func getInStr(lettersIdList []string) string {
	inStr := ""
	for index, letterId := range lettersIdList {
		if index == 0 {
			inStr = letterId
		} else {
			inStr = inStr + "," + letterId
		}
	}
	return inStr
}
