/*
Implement the dining philosopher’s problem with the following constraints/modifications.
 1. There should be 5 philosophers sharing chopsticks,
    with one chopstick between each adjacent pair of philosophers.
 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 5. The host allows no more than 2 philosophers to eat concurrently.
 6. Each philosopher is numbered, 1 through 5.
 7. When a philosopher starts eating (after it has obtained necessary locks)
    it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
 8. When a philosopher finishes eating (before it has released its locks)
    it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type CStick struct {
	sync.Mutex
}

type Philo struct {
	leftCStick, rightCStick *CStick
	id, count               int
}

func (p Philo) eat(c chan *Philo, wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {
		c <- &p
		if p.count < 3 {
			p.leftCStick.Lock()
			p.rightCStick.Lock()

			fmt.Println("starting to eat", p.id)
			p.count++
			fmt.Println("finishing eating", p.id)

			p.leftCStick.Unlock()
			p.rightCStick.Unlock()
			wg.Done()
		}
	}

}

func host(c chan *Philo, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan *Philo, 2)

	wg.Add(15)

	cSticks := make([]*CStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(CStick)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			leftCStick:  cSticks[i],
			rightCStick: cSticks[(i+1)%5],
			id: i+1,
			count: 0,
		}
	}

	go host(c, &wg)
	for i := 0; i < 5; i++ {
		go philos[i].eat(c, &wg)
	}
	wg.Wait()
}
