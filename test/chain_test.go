package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type student struct {
	Name string
	Age  int
}

type Broadcaster struct {
	subscribers []chan student
	response    []chan string
	mu          sync.Mutex
}

func (b *Broadcaster) AddSubscriber(subscriber chan student, response chan string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers = append(b.subscribers, subscriber)
	b.response = append(b.response, response)
}

func (b *Broadcaster) Send(msg student) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, sub := range b.subscribers {
		select {
		case sub <- msg:
		default:
			fmt.Println("Channel full, message dropped")
		}
	}
}

func (b *Broadcaster) CollectResponses() []string {
	b.mu.Lock()
	defer b.mu.Unlock()

	var responses []string
	for _, resChan := range b.response {
		select {
		case res := <-resChan:
			responses = append(responses, res)
		case <-time.After(2 * time.Second): // 超时处理
			responses = append(responses, "response timeout")
		}
	}

	return responses
}

func subscriber(id int, subChan <-chan student, resChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for student := range subChan {
		// 模拟处理
		time.Sleep(time.Second)
		resChan <- fmt.Sprintf("Subscriber %d received: %s, Age: %d", id, student.Name, student.Age)
	}
	close(resChan) // 处理完后关闭响应通道
}

func TestChan(t *testing.T) {
	b := &Broadcaster{}

	// 创建订阅者通道
	subChan1 := make(chan student, 1) // 使用缓冲通道来避免阻塞
	resChan1 := make(chan string)
	subChan2 := make(chan student, 1) // 使用缓冲通道来避免阻塞
	resChan2 := make(chan string)

	// 添加订阅者
	b.AddSubscriber(subChan1, resChan1)
	b.AddSubscriber(subChan2, resChan2)

	// 启动订阅者
	var wg sync.WaitGroup
	wg.Add(2)
	go subscriber(1, subChan1, resChan1, &wg)
	go subscriber(2, subChan2, resChan2, &wg)

	// 发送消息
	b.Send(student{Name: "Alice", Age: 21})
	b.Send(student{Name: "Bob", Age: 22})

	// 等待订阅者处理完
	time.Sleep(2 * time.Second) // 给订阅者一些时间来处理

	// 收集并打印响应
	responses := b.CollectResponses()
	for _, response := range responses {
		fmt.Println(response)
	}

	// 关闭通道
	close(subChan1)
	close(subChan2)

	// 等待所有订阅者完成
	wg.Wait()
}
