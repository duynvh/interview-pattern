package main

import "sync"

type singleton struct {
	counter int
}

var once sync.Once
var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton) // <- thread safe
	})

	return instance
}

func (s *singleton) Increase() {
	s.counter++
}

func (s *singleton) Get() int {
	return s.counter
}
