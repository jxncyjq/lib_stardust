package main

import (
	"sync"
)

// SingletonInterface 定义了单例接口
type SingletonInterface[T any] interface {
	Get() T
}

// Singleton 实现了 SingletonInterface
type Singleton[T any] struct {
	value T
	once  sync.Once
}

// Get 返回单例的值
func (s *Singleton[T]) Get() T {
	return s.value
}

// SingletonManager 用于管理不同类型的单例
type SingletonManager struct {
	instances map[string]interface{}
	onceMap   map[string]*sync.Once
	mu        sync.Mutex
}

// NewSingletonManager 创建一个 SingletonManager 实例
func NewSingletonManager() *SingletonManager {
	return &SingletonManager{
		instances: make(map[string]interface{}),
		onceMap:   make(map[string]*sync.Once),
	}
}

// GetInstance 获取单例实例
func (sm *SingletonManager) GetInstance[T any](key string, initFunc func() T) SingletonInterface[T] {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 检查是否已经创建了实例
	if inst, ok := sm.instances[key]; ok {
		return inst.(SingletonInterface[T])
	}

	// 如果没有创建，则初始化一次
	var once sync.Once
	sm.onceMap[key] = &once
	var instance SingletonInterface[T]
	once.Do(func() {
		instance = &Singleton[T]{value: initFunc()}
		sm.instances[key] = instance
	})

	return instance
}

func main() {
	manager := NewSingletonManager()

	// 示例：创建一个 int 类型的单例
	intInstance := manager.GetInstance("int", func() int {
		return 42
	})

	// 获取并打印单例值
	println(intInstance.Get())
}
