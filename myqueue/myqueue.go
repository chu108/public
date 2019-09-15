package myqueue

import (
	"runtime"
	"sync"

	"gopkg.in/eapache/queue.v1"
)

//MyQueue queue
type MyQueue struct {
	sync.Mutex
	popable *sync.Cond
	buffer  *queue.Queue
	closed  bool
}

//New 创建
func New() *MyQueue {
	ch := &MyQueue{
		buffer: queue.New(),
	}
	ch.popable = sync.NewCond(&ch.Mutex)
	return ch
}

//Pop 取出队列,（阻塞模式）
func (q *MyQueue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for buffer.Length() == 0 && !q.closed {
		c.Wait()
	}

	if q.closed { //已关闭
		return
	}

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
	}
	return
}

//试着取出队列（非阻塞模式）返回ok == false 表示空
func (q *MyQueue) TryPop() (v interface{}, ok bool) {
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		ok = true
	} else if q.closed {
		ok = true
	}

	return
}

// 插入队列，非阻塞
func (q *MyQueue) Push(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		q.buffer.Add(v)
		q.popable.Signal()
	}
}

// 获取队列长度
func (q *MyQueue) Len() int {
	return q.buffer.Length()
}

// Close MyQueue
// After close, Pop will return nil without block, and TryPop will return v=nil, ok=True
func (q *MyQueue) Close() {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		q.closed = true
		q.popable.Broadcast() //广播
	}
}

//Wait 等待队列消费完成
func (q *MyQueue) Wait() {
	for {
		if q.closed || q.buffer.Length() == 0 {
			break
		}

		runtime.Gosched() //出让时间片
	}
}
