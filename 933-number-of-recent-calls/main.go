package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.Ping(1))
	fmt.Println(c.Ping(100))
	fmt.Println(c.Ping(3001))
	fmt.Println(c.Ping(3002))
}

type RecentCounter struct {
	history []int
}


func Constructor() RecentCounter {
	return RecentCounter{
		history: make([]int, 0, 10),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.history = append(this.history, t)
	counter := 0
	limit := t-3000
	for i := len(this.history) - 1; i >= 0 ;i--{
		if this.history[i] < limit {
			break
		}
		counter++
	}
	return counter
}
