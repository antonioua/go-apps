package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func main() {
	wg := &sync.WaitGroup{}
	var once sync.Once

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if foundTreasure() {
				once.Do(markMissionCompleted) // Don't call function multiple times but once
			}
			wg.Done()
		}()
	}
	wg.Wait()

	checkMissionCompletion()
}

func markMissionCompleted() {
	missionCompleted = true
	fmt.Println("Marking mission completed!")
}
func foundTreasure() bool {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndNum := rnd.Intn(10)
	//fmt.Println("Random number: ", rndNum)
	return 0 == rndNum
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("Mission completed!")
	} else {
		fmt.Println("Mission failed!")
	}
}
