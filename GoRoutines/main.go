package main

import (
	"fmt"
	"time"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func goCallOdd(ctx context.Context, wg *sync.WaitGroup, i chan int, termChan chan os.Signal){
	defer wg.Done()
	go goSig(ctx, "Odd", wg, termChan)
	for v := range i {
		time.Sleep(1*time.Second)
		fmt.Println("Odd Value : ", v)
	}

}

func goCallEven(ctx context.Context, wg *sync.WaitGroup, i chan int, termChan chan os.Signal){
	defer wg.Done()
	go goSig(ctx, "Even", wg, termChan)
	for v := range i {
		time.Sleep(1*time.Second)
		fmt.Println("Even Value : ", v)
	}

}

func goSig(ctx context.Context, s string, wg *sync.WaitGroup, termChan chan os.Signal){
	for {
		select {
			case <-ctx.Done():
				fmt.Println("Graceuly Shuttingdown : ", s)
				wg.Wait()
				time.Sleep(5*time.Second)
				os.Exit(1)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	o := make(chan int)
	e := make(chan int)
	termChan := make(chan os.Signal)
	ctx := context.Background()
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-termChan
		fmt.Printf("system call:%+v\n", oscall)
		cancel()
	}()

	wg.Add(2)
	go goCallEven(ctx, &wg, e, termChan)
	go goCallOdd(ctx, &wg, o, termChan)
	for i:=0; i<=10; i++ {
		if i%2 == 0 {
			e<- i
		} else {
			o<- i
		}
	}
	time.Sleep(2*time.Second)
}