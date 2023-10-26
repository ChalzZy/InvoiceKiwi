package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	fmt.Println("Hello")
	cron()
}

func helloworld() {
	fmt.Println("Printing shit")
}

func cron() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SingletonMode()

	scheduler.Every(2).Seconds().Do(func() {
		helloworld()
	})

	scheduler.StartBlocking()
}
