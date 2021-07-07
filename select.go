package main

import (
	"fmt"
	"time"
)

func main()  {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(1)
		c2 <- 2
	}()

	fmt.Println("c2:", <-c2)
	select {
	case  a :=<- c1:
		fmt.Println("read to c1", a)
	case  t1 := <- time.After(3*time.Second):
		fmt.Println("timeout ...", t1)
	//case  b := <- c2:
		//fmt.Println("read to c2", b)

	//default:
	//	fmt.Println("default...")
	}

}
