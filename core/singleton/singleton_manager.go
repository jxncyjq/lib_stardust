package singleton

import "sync"

//// SingletonManager 用于管理不同类型的单例
//type SingletonManager struct {
//	instances map[string]interface{}
//	onceMap   map[string]*sync.Once
//	mu        sync.Mutex
//}
//
//// NewSingletonManager 创建一个 SingletonManager 实例
//func NewSingletonManager() *SingletonManager {
//	return &SingletonManager{
//		instances: make(map[string]interface{}),
//		onceMap:   make(map[string]*sync.Once),
//	}
//}
//
//// GetInstance 是一个独立的泛型函数，用于获取或创建单例实例
//func GetInstance[T any](sm *SingletonManager, key string, initFunc func() T) SingletonInterface[T] {
//	sm.mu.Lock()
//	defer sm.mu.Unlock()
//
//	// 检查是否已经创建了实例
//	if inst, ok := sm.instances[key]; ok {
//		return inst.(SingletonInterface[T])
//	}
//
//	// 如果没有创建，则初始化
//	instance := &Singleton[T]{value: initFunc()}
//	sm.instances[key] = instance
//
//	return instance
//}

// SingletonManager 管理所有单例实例
type SingletonManager struct {
	instances map[string]interface{}
	mu        sync.Mutex
}

// NewSingletonManager 创建一个新的 SingletonManager 实例
func NewSingletonManager() *SingletonManager {
	return &SingletonManager{
		instances: make(map[string]interface{}),
	}
}

// GetOrCreateInstance 获取或创建一个单例实例
func (sm *SingletonManager) GetOrCreateInstance(key string, initFunc func() interface{}) interface{} {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if instance, exists := sm.instances[key]; exists {
		return instance
	}

	instance := initFunc()
	sm.instances[key] = instance
	return instance
}
