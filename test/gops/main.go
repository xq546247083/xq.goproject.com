package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
		ConfigDir:       "Log",
		Addr:            ":10018",
	}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
