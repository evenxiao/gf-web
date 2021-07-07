package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gogf/gf"
)



const (
	WorkerNo = 5
	TaskNo = 10
)
var wg  sync.WaitGroup
func main() {
	fmt.Println("hello gf...", gf.VERSION)

	fmt.Println("werwerwe")

	log.Println(gf.AUTHORS)

	c1 := make(chan int, WorkerNo)

	for i:=1; i <= WorkerNo; i++ {

		wg.Add(1)
		go taskProcess(c1, i)

	}

	for t:= 1; t <= TaskNo; t++ {
		c1 <- t
	}
	close(c1)
	wg.Wait()
	fmt.Println("All work done!")

}
func taskProcess(c chan int, wokerNo int){

	defer  wg.Done()

	for v := range c{
		fmt.Printf("wokerNo %d is take process no:%d \n", wokerNo, v)
	}
	fmt.Printf("workerNo %d is got off work \n", wokerNo)
}


