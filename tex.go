package main

import (
	"fmt"
	"time"
)

func main() {
	var i = make(chan int)
	var j = make(chan int)

	go func() {
		for{
			select {
			case a:= <-i:{
				fmt.Println(a)
			}
			case a:=<-j:{
				fmt.Println(a)
			}
			}
		}

	}()
	go func() {i<-2}()
	go func() {j<-3}()
	time.Sleep(3*time.Second)
}
