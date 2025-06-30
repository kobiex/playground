package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for b.Loop() {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for b.Loop() {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
