package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg         sync.WaitGroup
		numberChan chan int
	)
	numberChan = make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("写入数据: %d\n", i)
			numberChan <- i
			time.Sleep(time.Duration(1) * time.Second)
		}
		close(numberChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for number := range numberChan {
			fmt.Printf("读取数据: %d\n", number)
		}
	}()

	var (
		numberBufferChan chan int
	)
	numberBufferChan = make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("写入数据Buffer: %d\n", i)
			numberBufferChan <- i
			time.Sleep(time.Duration(1) * time.Second)
		}
		close(numberBufferChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for number := range numberBufferChan {
			fmt.Printf("读取数据Buffer: %d\n", number)
		}
	}()

	wg.Wait()
}
