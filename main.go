package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

func main() {
	withSkipIfStillRunning()
	listen()
}

func withSkipIfStillRunning() {
	fmt.Println(time.Now().String() + " - Start App - WithSkipIfStillRunning")

	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
	))
	id, _ := c.AddFunc("@every 30s", task)
	c.Entry(id).Job.Run()

	go c.Start()
}

func task() {
	fmt.Println(time.Now().String() + " - Start Task")
	time.Sleep(40 * time.Second)
	fmt.Println(time.Now().String() + " - End Task")
}

func startImmediately() {
	fmt.Println(time.Now().String() + " - Start App - Start Immediately")
	c := cron.New()
	id, _ := c.AddFunc("@every 1m10s", task)
	//Utiliza o EntryId para invocar a tarefa imediatamente
	c.Entry(id).Job.Run()

	go c.Start()
}

func simpleCron() {
	fmt.Println(time.Now().String() + " - Start App - Version SimpleCron")
	c := cron.New()
	c.AddFunc("@every 1m10s", task)

	go c.Start()
}

func listen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println(time.Now().String() + " - Closed")
}
