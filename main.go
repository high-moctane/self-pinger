package main

import (
	"flag"
	"time"

	"github.com/high-moctane/self-pinger/pinger"
)

var url = flag.String("u", "0.0.0.0", "target url")
var dur = flag.Duration("d", 20*time.Minute, "ping duration")

func main() {
	p := pinger.Pinger{URL: *url, Duration: *dur}
	p.Run()
}
