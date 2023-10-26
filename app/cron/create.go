package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func RunCron() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SingletonMode()

	scheduler.Every(2).Seconds().Do(func() {
		fmt.Println("working!")
	})

	scheduler.StartBlocking()
}
