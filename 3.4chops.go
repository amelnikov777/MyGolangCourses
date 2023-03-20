// chops
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	hostCh_r := make(chan int, 2)     //Buffered request channel limiting simultaneous permissions for two goroutines
	hostCh_a := make(chan int)        //Permission channel
	phil_rand := rand.Intn(40)/10 + 5 //The first eating philosophe is random

	go host(hostCh_r, hostCh_a) //Starting host

	for i := 0; i < 5; i++ { //Starting eat method for 5 philosophes
		phil_eat := (phil_rand + i) % 5
		go philos[phil_eat].eat(phil_eat, hostCh_r, hostCh_a)
	}

	for pause := 0; pause < 1000000000; pause++ {
	} //Pause for finish all goroutines
}

//In this metod each philosof asks permission from host and starts eating is steacks ate free
func (p Philo) eat(phil int, hostCh_r, hostCh_a chan int) {
	phil_eat := phil + 1
	for t := 1; t < 4; t++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		hostCh_r <- phil_eat
		fmt.Println("Request from philosophe", phil_eat)
		permission := <-hostCh_a
		permission++
		fmt.Println("start eating phil.\t", phil_eat, "time", t)
		for eating_pause := 0; eating_pause < 1000000; eating_pause++ {
		} //Waiting for eating process
		fmt.Println("finish eating phil.\t", phil_eat, "time", t)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

//Host function check requests from "eat" method and gives permissions
func host(hostCh_r chan int, hostCh_a chan int) { //host
	for {
		req := <-hostCh_r
		for host_pause := 0; host_pause < 100; host_pause++ {
		}
		hostCh_a <- req

	}
}
