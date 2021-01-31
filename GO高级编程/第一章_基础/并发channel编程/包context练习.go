package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(goId int, ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
		case <-ctx.Done():
			fmt.Println(goId, ":\t", ctx.Err())
			return ctx.Err()
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, ctx, &wg)
	}

	// 3s后关闭ctx
	time.Sleep(1 * time.Second)
	cancel()

	wg.Wait()
}
