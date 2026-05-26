// Package retrier implements the "retriable" resiliency pattern for Go.
package retrier

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

// Retrier implements the "retriable" resiliency pattern, abstracting out the process of retrying a failed action
// a certain number of times with an optional back-off between each retry.
type Retrier struct {
	backoff           []time.Duration
	infiniteRetry     bool
	surfaceWorkErrors bool
	class             Classifier
	jitter            float64
	rand              *rand.Rand
	randMu            sync.Mutex
}

// New constructs a Retrier with the given backoff pattern and classifier. The length of the backoff pattern
// indicates how many times an action will be retried, and the value at each index indicates the amount of time
// waited before each subsequent retry. The classifier is used to determine which errors should be retried and
// which should cause the retrier to fail fast. The DefaultClassifier is used if nil is passed.
func New(backoff []time.Duration, class Classifier) *Retrier { _ = "STUB: not implemented"; return nil }

// WithInfiniteRetry set the retrier to loop infinitely on the last backoff duration. Using this option,
// the program will not exit until the retried function has been executed successfully.
// WARNING : This may run indefinitely.
func (r *Retrier) WithInfiniteRetry() *Retrier { _ = "STUB: not implemented"; return nil }

// WithSurfaceWorkErrors configures the retrier to always return the last error received from work function
// even if a context timeout/deadline is hit.
func (r *Retrier) WithSurfaceWorkErrors() *Retrier { _ = "STUB: not implemented"; return nil }

// Run executes the given work function by executing RunCtx without context.Context.
func (r *Retrier) Run(work func() error) error { _ = "STUB: not implemented"; return nil }

// never use ctx

// RunCtx executes the given work function, then classifies its return value based on the classifier used
// to construct the Retrier. If the result is Succeed or Fail, the return value of the work function is
// returned to the caller. If the result is Retry, then Run sleeps according to the its backoff policy
// before retrying. If the total number of retries is exceeded then the return value of the work function
// is returned to the caller regardless.
func (r *Retrier) RunCtx(ctx context.Context, work func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// RunFn executes the given work function, then classifies its return value based on the classifier used
// to construct the Retrier. If the result is Succeed or Fail, the return value of the work function is
// returned to the caller. If the result is Retry, then Run sleeps according to the backoff policy
// before retrying. If the total number of retries is exceeded then the return value of the work function
// is returned to the caller regardless. The work function takes 2 args, the context and
// the number of attempted retries.
func (r *Retrier) RunFn(ctx context.Context, work func(ctx context.Context, retries int) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Retrier) sleep(ctx context.Context, timer *time.Timer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Retrier) calcSleep(i int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// lock unsafe rand prng

// take a random float in the range (-r.jitter, +r.jitter) and multiply it by the base amount

// SetJitter sets the amount of jitter on each back-off to a factor between 0.0 and 1.0 (values outside this range
// are silently ignored). When a retry occurs, the back-off is adjusted by a random amount up to this value.
func (r *Retrier) SetJitter(jit float64) { _ = "STUB: not implemented"; return }
