package main

import (
	"fmt"
	"time"
	"sync"
)


func main(){
	lock :=make(chan bool,1)
	var wg sync.WaitGroup
	for i:=1;i<7;i++{
		wg.Add(1)
		go func(i int,lock chan bool){
			fmt.Printf("worker %d wants the lock \n", i)
			lock <-true
			fmt.Printf("%d has the lock\n", i)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%d is releasing the lock\n", i)
			<-lock
			wg.Done()
		}(i,lock)
	}
	wg.Wait()
	//time.Sleep(2000*time.Millisecond)
	fmt.Printf("Compressed %d files\n", 0)
}


