/*
 * @Descripttion:
 * @Version: 1.0
 * @Author: ouyangtianpeng ouyangtianpeng@jxcc.com
 * @Date: 2024-10-12 08:32:59
 */
package basic

import "fmt"

func TestArray() {
	var interfaceArr [3]interface{}
	for i, v := range interfaceArr {
		fmt.Printf("interfaceArr[%d] = %v", i, v)
	}
}
