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
	for {
		resp, err := http.Get(p.URL)
		if err != nil {
			log.Println("pinger: ", err)
		}
		log.Println("pinger: ping!")
		if err = resp.Body.Close(); err != nil {
			log.Println("pinger: ", err)
		}
		time.Sleep(p.Duration)
	}
}
