// 5 philosophes eat rice with 5 chopsticks
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

	hostCh_r := make(chan int, 2)          //Buffered request channel limiting simultaneous permissions for two goroutines
	hostCh_a := make(chan int)             //Permission channel
	phil_rand := (rand.Intn(40)+10)/10 + 5 //The first eating philosophe is random

	go host(hostCh_r, hostCh_a) //Starting host

	for i := 0; i < 5; i++ { //Starting eat method for 5 philosophes
		phil_eat := (phil_rand + i) % 5
		go philos[phil_eat].eat(phil_eat, hostCh_r, hostCh_a)
	}
	pause(1000000000) //Pause for finishing all goroutines
}

//In this method each philosophe asks permission from host and starts eating if chopsticks are free
func (p Philo) eat(phil int, hostCh_r, hostCh_a chan int) {
	phil_eat := phil + 1
	for t := 1; t < 4; t++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		hostCh_r <- phil_eat
		fmt.Println("Host request from philosop.\t", phil_eat)
		permission := <-hostCh_a
		permission++
		fmt.Println("\tstarting to eat\t", phil_eat, "time", t)
		pause(10000000) //Waiting for eating process
		fmt.Println("\tfinishing eating\t", phil_eat, "time", t)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

//Host function checks requests from "eat" method and gives permissions
func host(hostCh_r chan int, hostCh_a chan int) { //host
	for {
		req := <-hostCh_r
		pause(1000)
		hostCh_a <- req
	}
}

func pause(time int) {
	for pause := 0; pause < time; pause++ {
	}
}
