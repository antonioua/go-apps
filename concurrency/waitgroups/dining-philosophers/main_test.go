package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	sleepTime = time.Second * 0
	eatTime = time.Second * 0
	thinkTime = time.Second * 0

	for i := 0; i < 100; i++ {
		main()
		if len(personsFinished) != 5 {
			t.Errorf("Wrong number of finished persons")
		}

		personsFinished = []string{}
	}

}
