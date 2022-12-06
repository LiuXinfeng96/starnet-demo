package services

import (
	"sync"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/golang/groupcache/lru"
)

type SdkPool struct {
	cache *lru.Cache

	mutex sync.Mutex
}

func NewSdkPool(maxEntries int) *SdkPool {
	var sdkPool SdkPool
	sdkPool.cache = lru.New(maxEntries)
	sdkPool.cache.OnEvicted = func(key lru.Key, value interface{}) {
		client, ok := value.(sdk.ChainClient)
		if !ok {
			return
		}
		client.Stop()
	}

	return &sdkPool
}

func (s *SdkPool) Add(key string, client *sdk.ChainClient) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Add(key, client)
}

func (s *SdkPool) Get(key string) (client *sdk.ChainClient, ok bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	v, ok := s.cache.Get(key)

	client, ok = v.(*sdk.ChainClient)

	return
}

func (s *SdkPool) Remove(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Remove(key)
}

func (s *SdkPool) RemoveOldest() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.RemoveOldest()
}

func (s *SdkPool) Len() int {
	return s.cache.Len()
}

func (s *SdkPool) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Clear()
}
