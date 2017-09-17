package main

import (
	"flag"
	"time"

	"github.com/high-moctane/self-pinger/pinger"
)

var url = flag.String("u", "http://localhost", "target url")
var dur = flag.Duration("d", 20*time.Minute, "ping duration")

func main() {
	flag.Parse()
	p := pinger.Pinger{URL: *url, Duration: *dur}
	p.Run()
}
