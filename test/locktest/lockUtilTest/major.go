package lockUtilTest

import "sync"

var (
	// 针对锁map的锁
	localLock sync.RWMutex

	// 锁map
	lockMap = make(map[int]*sync.Mutex)
)

// GetLock 获取锁
//@key 关键字
//@return 返回值
func GetLock(key int) *sync.Mutex {
	// 是否存在锁
	existsFunc := func() (*sync.Mutex, bool) {
		localLock.RLock()
		defer localLock.RUnlock()

		result, exists := lockMap[key]
		return result, exists
	}

	// 如果存在，直接返回
	if result, exists := existsFunc(); exists {
		return result
	}

	// 不存在，添加锁
	localLock.Lock()
	defer localLock.Unlock()
	lockMap[key] = new(sync.Mutex)

	return lockMap[key]
}
