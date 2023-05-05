

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/job"
)

func main() {
	ch := make(chan os.Signal, 1)
	a := job.New(job.WithSeconds(), job.WithChain(job.Recover(g_log.V(1))), job.WithLogger(g_log.V(1)))
	a.AddFunc("*/10 * * * * *", func() { log.Printf("Every 10 second") })
	a.AddFunc("@hourly", func() { log.Printf("Every hour") })
	a.AddFunc("@every 1h30m", func() { log.Printf("Every hour thirty") })
	a.AddFunc("@every 1m", func() {
		log.Printf("Every minute")
	})

	a.Start()
	log.Printf("Waiting for stop")
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

}
