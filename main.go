/*
 * @Descripttion:
 * @Version: 1.0
 * @Author: ouyangtianpeng ouyangtianpeng@jxcc.com
 * @Date: 2024-10-11 15:18:40
 */
package main

import (
	"fmt"

	"example.com/test/src/basic"
)

func main() {
	// basic.TestChannel()
	// basic.TestArray()
	basic.TestCircularQueue()
}

func intToHexStr(ret int) string {
	hexStr := fmt.Sprintf("0x%X", ret)
	return hexStr
}
