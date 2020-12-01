package utils

import (
	"sync"
)

type LimitClick struct {
	sync.Mutex
	userMap map[string]int64
}

func InitLimitClick() *LimitClick {
	return &LimitClick{
		userMap: make(map[string]int64),
	}
}

func (c *LimitClick) Exists(userID string) bool {
	c.Lock()
	defer c.Unlock()

	_, ok := c.userMap[userID]
	if !ok {
		c.userMap[userID] = 1
	}
	return ok
}

func (c *LimitClick) Delete(userID string) {
	c.Lock()
	defer c.Unlock()

	delete(c.userMap, userID)
}
