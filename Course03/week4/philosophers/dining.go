package main

import (
	"fmt"
	"sync"
)

const (
	philo   = 5
	eatNum  = 3
	eatConc = 2
)

var (
	wg sync.WaitGroup
)

type Chops struct {
	sync.Mutex
}

type Host struct {
	eatAllowed chan int
}

func (host *Host) allowToEat(ph *Philosopher) {
	go func() {
		<-host.eatAllowed
		ph.eat()
	}()

	go func() {
		host.eatAllowed <- 1
	}()
}

type Philosopher struct {
	leftCS  *Chops
	rightCS *Chops
	Id      int
}

func (p *Philosopher) eat() {
	defer wg.Done()

	// Pick chopstics
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("starting to eat %d\n", p.Id)
	for i := 1; i <= eatNum; i++ {
		fmt.Printf(" philosopher %d: serving %d\n", p.Id, i)
	}

	fmt.Printf("finishing eating %d\n", p.Id)
	p.rightCS.Unlock()
	p.leftCS.Unlock()

}

func main() {

	// init 5 chopstics
	CSticks := make([]*Chops, philo)
	for i := 0; i < philo; i++ {
		CSticks[i] = new(Chops)
	}

	// init 5 philosopehrs
	philos := make([]*Philosopher, philo)
	for i := 0; i < philo; i++ {
		// last element i=4 will take first element (i+1)%5 = 0
		philos[i] = &Philosopher{
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%philo],
			Id:      i + 1,
		}
	}

	// init host with 2 available concurrent servings
	host := &Host{eatAllowed: make(chan int, eatConc)}
	for i := 0; i <= eatConc; i++ {
		go func() { host.eatAllowed <- 1 }()
	}

	// run the process of eating
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go host.allowToEat(philos[i])
	}

	wg.Wait()
}
