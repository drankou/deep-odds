package utils

import (
	"testing"
	"time"
)

func TestRateLimiter_Init(t *testing.T) {
	limiter := &RateLimiter{}
	limiter.Init(time.Second)

	if limiter.rate != time.Second{
		t.Fatal("limiter rate is not set")
	}

	if limiter.Mutex == nil{
		t.Fatal("rate limiter mutex is nil")
	}
}

func TestRateLimiter_RateBlock(t *testing.T) {
	limiter := RateLimiter{}
	rate := time.Second
	limiter.Init(rate)

	for i:=0; i<3; i++{
		before := time.Now()
		limiter.RateBlock()
		if time.Since(before) < rate{
			t.Fatal("rate is higher than limit")
		}
	}
}
