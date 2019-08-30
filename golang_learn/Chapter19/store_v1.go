package mypkg

import "sync"

type URLStore struct {
	urls map[string]string
	mu sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls:make(map[string]string)}
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.UnRLock
	return s.urls.map[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.WLock()
	s.mu.UnWLock()
	_, present := s.urls[key]
	if present{
		return false
	}
	s.ruls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.UnRLock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for{
		key := genKey(url)
		if s.Set(key, url){
			return key
		}
	}
	return ""
}

func genKey(n int) string {
	keyChar := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}
