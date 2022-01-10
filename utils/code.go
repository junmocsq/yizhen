package utils

import "fmt"

var msgMap = map[int]string{
	200:  "SUCCESS",
	4000: "参数 %s 丢失",
}

func Msg(code int, params ...interface{}) string {
	if _, ok := msgMap[code]; !ok {
		return fmt.Sprintf("code %d 无定义", code)
	}
	return fmt.Sprintf(msgMap[code], params...)
}
