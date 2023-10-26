package main

import (
	"cron"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	cron.RunCron()
}
