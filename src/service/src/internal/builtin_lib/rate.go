package builtin_lib

import (
	"math"
	"sync"
	"time"
)

type Limit float64

const Inf = Limit(math.MaxFloat64)

func Every(interval time.Duration) Limit {
	if interval <= 0 {
		return Inf
	}
	return 1 / Limit(interval.Seconds())
}

type Limiter struct {
	mutex     sync.Mutex
	limit     Limit
	burst     int
	tokens    float64
	last      time.Time
	lastEvent time.Time
}

func NewLimiter(limit Limit, burst int) *Limiter {
	return &Limiter{
		limit: limit,
		burst: burst,
	}
}

func (limiter *Limiter) Allow() bool {
	return limiter.reserve(time.Now(), 1, 0).status
}

type Reservation struct {
	status bool
	tokens int

	timeToAct time.Time
	limiter   *Limiter
	limit     Limit
}

const InfDuration = time.Duration(1<<63 - 1)

func (limiter *Limiter) reserve(now time.Time, nTokens int, maxFutureReserve time.Duration) Reservation {
	limiter.mutex.Lock()
	defer limiter.mutex.Unlock()

	if limiter.limit == Inf {
		return Reservation{
			status:    true,
			limiter:   limiter,
			tokens:    nTokens,
			timeToAct: now,
		}
	} else if limiter.limit == 0 {
		var status bool
		if limiter.burst >= nTokens {
			status = true
			limiter.burst -= nTokens
		}
		return Reservation{
			status:    status,
			limiter:   limiter,
			tokens:    limiter.burst,
			timeToAct: now,
		}
	}

	now, last, tokens := limiter.calculate(now)

	tokens -= float64(nTokens)

	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = limiter.limit.durationFromTokens(-tokens)
	}

	status := nTokens <= limiter.burst && waitDuration <= maxFutureReserve

	reservation := Reservation{
		status:  status,
		limiter: limiter,
		limit:   limiter.limit,
	}
	if status {
		reservation.tokens = nTokens
		reservation.timeToAct = now.Add(waitDuration)
	}

	if status {
		limiter.last = now
		limiter.tokens = tokens
		limiter.lastEvent = reservation.timeToAct
	} else {
		limiter.last = last
	}

	return reservation
}

func (limiter *Limiter) calculate(now time.Time) (newNow time.Time, newLast time.Time, newTokens float64) {
	last := limiter.last
	if now.Before(last) {
		last = now
	}

	elapsed := now.Sub(last)
	delta := limiter.limit.tokensFromDuration(elapsed)
	tokens := limiter.tokens + delta
	if burst := float64(limiter.burst); tokens > burst {
		tokens = burst
	}
	return now, last, tokens
}

func (limit Limit) durationFromTokens(tokens float64) time.Duration {
	if limit <= 0 {
		return InfDuration
	}
	seconds := tokens / float64(limit)
	return time.Duration(float64(time.Second) * seconds)
}

func (limit Limit) tokensFromDuration(d time.Duration) float64 {
	if limit <= 0 {
		return 0
	}
	return d.Seconds() * float64(limit)
}
