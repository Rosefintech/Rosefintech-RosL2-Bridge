/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/5/11 14:44
*/

package tools

import "strings"

func CancelTargetStr(str, substr string) string {
	var (
		resStr string
		index  int
	)
	for {
		if index = strings.Index(str, substr); index == -1 {
			resStr += str
			break
		}
		resStr += str[:index]
		str = str[index+len(substr):]
	}
	return resStr
}
