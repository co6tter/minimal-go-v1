package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	fmt.Println("Something done!")
	wg.Done() // WaitGroupは自動的にデリファレンスしてくれる
}

func double(num int, ch chan int) {
	time.Sleep(time.Second * 1)
	ch <- num * 2
}

func getWeather1a(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "Sunny"
}

func getWeather1b(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "A bit sunny"
}

func getWeather1c(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "Almost  sunny"
}

func increment(count *int, mu *sync.Mutex) {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		*count++
		mu.Unlock()
	}
}

func main() {
	fmt.Println("Start!")

	var wg sync.WaitGroup
	wg.Add(3)
	start := time.Now()
	go doSomething(&wg)
	go doSomething(&wg)
	go doSomething(&wg)

	wg.Wait()
	fmt.Println(time.Since(start))

	// var ch chan int = make(chan int)
	ch := make(chan int)

	go double(10, ch)
	go double(20, ch)
	go double(30, ch)

	// result1 := <-ch
	// result2 := <-ch
	// result3 := <-ch
	// fmt.Println(result1 + result2 + result3)
	fmt.Println(<-ch + <-ch + <-ch)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go double(10, ch1)
	go double(20, ch2)
	go double(30, ch3)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)

	ch1a := make(chan string)
	ch1b := make(chan string)
	ch1c := make(chan string)

	go getWeather1a(ch1a)
	go getWeather1b(ch1b)
	go getWeather1c(ch1c)

	select {
	case weather := <-ch1a:
		fmt.Println(weather)
	case weather := <-ch1b:
		fmt.Println(weather)
	case weather := <-ch1c:
		fmt.Println(weather)
	case <-time.After(time.Second * 2):
		fmt.Println("request timeout")
	}

	// レースコンディションを解決するためにMutex
	var mu sync.Mutex
	count := 0

	go increment(&count, &mu)
	go increment(&count, &mu)
	go increment(&count, &mu)
	go increment(&count, &mu)
	go increment(&count, &mu)

	time.Sleep(time.Second * 2)
	fmt.Println(count)
	fmt.Println("Finish!")
}
