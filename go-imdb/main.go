package main

import (
	"context"
	"fmt"
	"github.com/antonioua/go-imdb/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // SIGINT = Ctrl + C

	fmt.Println("Press CTRL + C to interrupt main")

	args := os.Args
	cf := config.ParseConfig(args[1:])
	fmt.Println(cf)

	go doWork()

	maxRunTime, err := time.ParseDuration(cf.MaxRunTime)
	if err != nil {
		log.Fatal("Unable to parse `maxRequests` argument", err)
	}
	fmt.Println(maxRunTime)

	<-sigCh
	_, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer handleTermination(cancel)

	fmt.Println("Done")
}

func handleTermination(cancel context.CancelFunc) {
	// Todo: release all resources
	// finish goroutines
	fmt.Println("bye.")
	cancel()
}

func doWork() {
	// readfile
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 2)
	}
}
