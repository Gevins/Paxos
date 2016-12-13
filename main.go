package main

import (
	"fmt"
	// "time"
)

func main() {
	// timer1 := time.NewTimer(time.Second * 2)
	// go func() {
	// 	<-timer1.C
	// 	fmt.Println("Timer 1 expired")
	// }()

	// timer2 := time.NewTimer(time.Millisecond * 100)
	// go func() {
	// 	<-timer2.C
	// 	fmt.Println("Timer 2 expired")
	// }()
	// c2 := make(chan int, 3)

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		c2 <- i
	// 	}
	// 	// close(c2)
	// }()
	// t1 := time.Now()
	// count := 1
	// for i := range c2 {
	// 	fmt.Println(i)
	// 	// fmt.Println(t2.Sub(t1).Seconds())
	// 	if time.Now().Sub(t1).Nanoseconds()/1000 > 2 {
	// 		fmt.Println("ok")
	// 		break
	// 	}
	// 	// fmt.Println(t2.Sub(t1).Seconds())

	// 	if count > 5 {
	// 		break
	// 	}
	// }
	// select {
	// case res := <-c2:
	// 	fmt.Println(res)
	// case <-time.After(time.Second * 10):
	// 	fmt.Println("timeout 2")
	// }
	ok, str := test(10)
	fmt.Println(ok, str)
}

func test(len int) (bool, []int) {
	str := make([]int, len)
	str[0] = 1
	return true, str
}
