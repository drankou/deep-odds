package utils

import (
	"sync"
	"time"
)

type RateLimiter struct {
	*sync.Mutex
	rate time.Duration
	last time.Time
}

func (r *RateLimiter) Init(rate time.Duration) {
	r.Mutex = &sync.Mutex{}
	r.last = time.Now()
	r.rate = rate
}

func (r *RateLimiter) RateBlock() {
	r.Lock()
	defer r.Unlock()

	if time.Since(r.last) < r.rate {
		<-time.After(r.last.Add(r.rate).Sub(time.Now()))
	}
	r.last = time.Now()
}
