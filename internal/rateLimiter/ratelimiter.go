package ratelimiter

import "time"

type Limiter interface {
	Allow(ip string) (bool, time.Duration)
}

type Config struct {
	RequestsPerTImeFrame int
	TimeFrame            time.Duration
	Enabled               bool
}