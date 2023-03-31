package index

import "sync"

const maxBuckets = 12

type sumCache struct {
	buckets [maxBuckets]*sums
}

type sums struct {
	mux   *sync.RWMutex
	cache map[string]string
}

func newSumCache() *sumCache {
	buckets := [maxBuckets]*sums{}
	for i := range buckets {
		buckets[i] = newSums()
	}

	return &sumCache{
		buckets: buckets,
	}
}

func newSums() *sums {
	return &sums{
		mux:   &sync.RWMutex{},
		cache: map[string]string{},
	}
}

func (s *sumCache) set(key string, val string) {
	s.buckets[getKeyHash(key)].set(key, val)
}

func (s *sumCache) get(key string) string {
	return s.buckets[getKeyHash(key)].get(key)
}

func (s *sums) set(key string, val string) {
	s.mux.Lock()
	s.cache[key] = val
	s.mux.Unlock()
}

func (s *sums) get(key string) string {
	s.mux.RLock()
	sum := s.cache[key]
	s.mux.RUnlock()

	return sum
}
