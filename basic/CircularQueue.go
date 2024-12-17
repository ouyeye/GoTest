/*
 * @Descripttion:
 * @Version: 1.0
 * @Author: ouyangtianpeng ouyangtianpeng@jxcc.com
 * @Date: 2024-10-23 10:37:28
 */
package basic

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func TestCircularQueue() {
	cq := NewCircularQueue(30)
	var flag bool
	go func(cq *circularQueue) {
		for {
			time.Sleep(70 * time.Millisecond)
			randomNumber := rand.Intn(1000000)
			// 将随机数字转换为字符串
			numberString := fmt.Sprintf("%d", randomNumber)
			// 转换为[]byte
			numberBytes := []byte(numberString)
			cq.push(numberBytes)
			if flag {
				return
			}
		}
	}(cq)

	// go func(cq *circularQueue) {
	// 	for {
	// 		time.Sleep(120 * time.Millisecond)
	// 		cq.pop()
	// 		if flag {
	// 			return
	// 		}
	// 	}
	// }(cq)

	go func(cq *circularQueue) {
		for {
			time.Sleep(200 * time.Millisecond)
			count := uint(rand.Intn(8)) + 1
			cq.popByCount(count)
			if flag {
				return
			}
		}
	}(cq)

	time.Sleep(3000 * time.Millisecond)
	flag = true
}

type CircularQueue interface {
	push(byteData []byte)
	pop() ([]byte, bool)
	// 尝试一次出队count个元素，返回实际的出队数据以及真实出队的元素个数
	popByCount(count uint) ([][]byte, uint)
}

var _ CircularQueue = (*circularQueue)(nil)

type circularQueue struct {
	dataArray [][]byte
	front     uint
	rear      uint
	capacity  uint
	mutex     sync.Mutex
}

func NewCircularQueue(capacity uint) *circularQueue {
	return &circularQueue{
		dataArray: make([][]byte, capacity),
		capacity:  capacity,
		front:     0,
		rear:      0,
	}
}

func (q *circularQueue) push(byteData []byte) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	// 队列已满，丢弃最旧的数据
	if (q.rear+1)%q.capacity == q.front {
		fmt.Println(" 丢弃数据 data = ", q.dataArray[q.front])
		q.front = (q.front + 1) % q.capacity
	}
	q.dataArray[q.rear] = byteData
	q.rear = (q.rear + 1) % q.capacity
	fmt.Println(" push data = ", byteData)
}

func (q *circularQueue) pop() ([]byte, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	// 队列是否为空
	if q.front == q.rear {
		fmt.Println(" pop data = nil")
		return nil, false
	}
	byteData := q.dataArray[q.front]
	q.front = (q.front + 1) % q.capacity
	fmt.Println(" pop data = ", byteData, "ok = ", true)
	return byteData, true
}

func (q *circularQueue) popByCount(count uint) ([][]byte, uint) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var realCount uint
	var resultArray [][]byte
	count2 := count
	if count == 0 || q.front == q.rear {
		fmt.Println(" popByCount datas = ", resultArray, "realcount = ", realCount, " count = ", count2)
		return nil, 0
	}
	for count > 0 {
		if q.front == q.rear {
			break
		}
		resultArray = append(resultArray, q.dataArray[q.front])
		q.front = (q.front + 1) % q.capacity
		count = count - 1
		realCount = realCount + 1
	}
	fmt.Println(" popByCount datas = ", resultArray, "realcount = ", realCount, " count = ", count2)
	return resultArray, realCount
}
