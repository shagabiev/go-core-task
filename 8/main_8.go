package main

import "time"

type CustomWaitGroup struct {
	addChan chan int
	signal  chan struct{}
	counter int
}

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{
		addChan: make(chan int),
		signal:  make(chan struct{}, 1), // буфер 1, чтобы не блокироваться
		counter: 0,
	}
	go wg.run()
	return wg
}

// менеджер счётчика
func (wg *CustomWaitGroup) run() {
	for delta := range wg.addChan {
		wg.counter += delta
		if wg.counter == 0 {
			// сигнализируем о завершении
			select {
			case wg.signal <- struct{}{}:
			default: // если уже есть сигнал — не блокируем
			}
		}
	}
}

func (wg *CustomWaitGroup) Add(n int) {
	wg.addChan <- n
}

func (wg *CustomWaitGroup) Done() {
	wg.addChan <- -1
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.signal
}

func main() {
	wg := NewCustomWaitGroup()
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			println("Goroutine", id, "done")
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	println("All goroutines finished")
}
