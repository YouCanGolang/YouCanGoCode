package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg               sync.WaitGroup
		numberChan       chan int
		numberBufferChan chan int
	)
	numberChan = make(chan int)
	numberBufferChan = make(chan int, 5)

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
		for i := 100; i < 110; i++ {
			fmt.Printf("写入数据Buffer: %d\n", i)
			numberBufferChan <- i
			time.Sleep(time.Duration(1) * time.Second)
		}
		close(numberBufferChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case number1, ok := <-numberChan:
				if !ok {
					numberChan = nil
				} else {
					fmt.Printf("读取数据: %d\n", number1)
				}
			case number2, ok := <-numberBufferChan:
				if !ok {
					numberBufferChan = nil
				} else {
					fmt.Printf("读取数据Buffer: %d\n", number2)
				}
			}
			if numberChan == nil && numberBufferChan == nil {
				break
			}
		}
	}()

	wg.Wait()
}
