package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func StudentRegister(name string, age int) *Student {
//	return &Student{
//		Name: name,
//		Age:  age,
//	}
//}
//
//var mu sync.Mutex
//var chain string

func worker(ctx context.Context) {
FOR:
	for {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				fmt.Println("明天记得来加班")
				break FOR
			}
		default:
			fmt.Println("认真学习,天天向上")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var (
		wg sync.WaitGroup
	)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()

	//time.Sleep(3 * time.Second)
}
