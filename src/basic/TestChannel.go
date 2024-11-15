/*
 * @Descripttion:
 * @Version: 1.0
 * @Author: ouyangtianpeng ouyangtianpeng@jxcc.com
 * @Date: 2024-10-11 15:19:34
 */
package basic

import (
	"fmt"
	"runtime"
	"time"
)

func TestChannel() {
	// TestTime()
	GeneratePrimeNumber()
}

// 1.管道和定时器一起使用
func TestTime() {
	ch := make(chan int)
	ch4 := make(chan int)

	go func(c chan int) {
		count1 := 0
		for {
			time.Sleep(time.Second * 1)

			count1 = count1 + 1
			fmt.Println("in count1 = ", count1)
			c <- 1
		}
	}(ch)
	go func(c chan int) {
		count4 := 0
		for {
			time.Sleep(time.Second * 1)

			count4 = count4 + 1
			fmt.Println("in count4 = ", count4)
			c <- 4
		}
	}(ch4)
	count1 := 0
	count4 := 0
	for {
		time.Sleep(time.Millisecond * 2000)
		select {
		case <-ch:
			count1 = count1 + 1
			fmt.Println("out count1 = ", count1)
		case <-ch4:
			count4 = count4 + 1
			fmt.Println("out count4 = ", count4)
		}
	}
}

// 2. 管道实现素数筛
func GeneratePrimeNumber() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	fmt.Println("oytp goroutine count = ", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
		fmt.Println("oytp goroutine count = ", runtime.NumGoroutine())
	}
}

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
