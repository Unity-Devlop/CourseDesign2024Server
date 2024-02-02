package game

import (
	"fmt"
	"sync"
	"testing"
)

func Test_BoardCast(t *testing.T) {
	bs := NewBroadcastService[int]()
	chBroadcast := bs.Run()
	chA := bs.Listen()
	chB := bs.Listen()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range chA {
			fmt.Println("A", v)
		}
		wg.Done()
	}()
	go func() {
		for v := range chB {
			fmt.Println("B", v)
		}
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		chBroadcast <- i
	}
	bs.UnListen(chA)
	for i := 3; i < 6; i++ {
		chBroadcast <- i
	}
	close(chBroadcast)
	wg.Wait()
}
