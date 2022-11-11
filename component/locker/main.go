package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type locker struct {
	lck *sync.RWMutex
	store map[string]*sync.RWMutex
}

func (l *locker) Lock(key string) error {
	l.lck.Lock()

	if k, ok := l.store[key]; ok {
		l.lck.Unlock()
		k.Lock()
		return nil
	}

	newLock := new(sync.RWMutex)
	l.store[key] = newLock
	l.lck.Unlock()

	newLock.Lock()
	return nil
}

func (l *locker) Unlock(key string) error {
	l.lck.Lock()
	defer l.lck.Unlock()

	if k, ok := l.store[key]; ok {
		k.Unlock()
	}

	return nil
}

func main() {
	myLocker := locker{
		lck: new(sync.RWMutex),
		store: make(map[string]*sync.RWMutex),
	}

	counter := 0

	for k := 1; k <= 5; k++ {
		go func(thisK int) {
			for j := 1; j <= 10000; j++ {
				_ = myLocker.Lock(fmt.Sprintf("%d", 1))
				counter++
				_ = myLocker.Unlock(fmt.Sprintf("%d", 1))
			}
		}(k)
	}

	time.Sleep(time.Second * 4)
	log.Println(counter)
}