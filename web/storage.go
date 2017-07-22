package web

import (
	"github.com/bingbaba/hhsecret"
	"sync"
)

var (
	mu          sync.RWMutex
	STORAGE_MEM = make(map[string]interface{})
)

func GetClientByUser(username string) (*hhsecret.Client, bool) {
	mu.RLock()
	defer mu.RUnlock()

	v, found := STORAGE_MEM[username+".client"]
	if !found {
		return nil, false
	} else {
		return v.(*hhsecret.Client), true
	}
}

func SaveClient(username string, client *hhsecret.Client) {
	mu.Lock()
	defer mu.Unlock()
	STORAGE_MEM[username+".client"] = client
}
