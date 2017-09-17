package pinger

import (
	"log"
	"net/http"
	"time"
)

// Pinger fetches url with specified interval.
type Pinger struct {
	URL      string
	Duration time.Duration
}

// Run starts pinging.
func (p *Pinger) Run() {
	c := time.Tick(p.Duration)
	for range c {
		resp, err := http.Get(p.URL)
		if err != nil {
			log.Println("pinger: ", err)
			continue
		}
		log.Println("pinger: ping!")
		if err = resp.Body.Close(); err != nil {
			log.Println("pinger: ", err)
		}
	}
}
