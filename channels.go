package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum //send sum to c
}

func main() {
	s := []int{7, 2, 8, -1, 4, 9}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c

	fmt.Println("Begin") //每次x,y的值不会一样

	fmt.Println(x, y, x+y) //每次x,y的值不会一样

	fmt.Println("End")
}
