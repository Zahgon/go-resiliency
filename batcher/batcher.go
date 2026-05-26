// Package batcher implements the batching resiliency pattern for Go.
package batcher

import (
	"sync"
	"time"
)

type work struct {
	param  interface{}
	future chan error
}

// Batcher implements the batching resiliency pattern
type Batcher struct {
	timeout   time.Duration
	prefilter func(interface{}) error

	lock         sync.Mutex
	submit       chan *work
	doWork       func([]interface{}) error
	batchCounter sync.WaitGroup
	flushTimer   *time.Timer
}

// New constructs a new batcher that will batch all calls to Run that occur within
// `timeout` time before calling doWork just once for the entire batch. The doWork
// function must be safe to run concurrently with itself as this may occur, especially
// when the doWork function is slow, or the timeout is small.
func New(timeout time.Duration, doWork func([]interface{}) error) *Batcher {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the work function with the given parameter, possibly
// including it in a batch with other calls to Run that occur within the
// specified timeout. It is safe to call Run concurrently on the same batcher.
func (b *Batcher) Run(param interface{}) error { _ = "STUB: not implemented"; return nil }

// Prefilter specifies an optional function that can be used to run initial checks on parameters
// passed to Run before being added to the batch. If the prefilter returns a non-nil error,
// that error is returned immediately from Run and the batcher is not invoked. A prefilter
// cannot safely be specified for a batcher if Run has already been invoked. The filter function
// specified must be concurrency-safe.
func (b *Batcher) Prefilter(filter func(interface{}) error) { _ = "STUB: not implemented"; return }

func (b *Batcher) submitWork(w *work) { _ = "STUB: not implemented"; return }

// kick off a new batch if needed

// then add this work to the current batch

func (b *Batcher) batch(input <-chan *work) { _ = "STUB: not implemented"; return }

// Shutdown flushes and executes any pending batches. If wait is true, it also waits for the pending batches
// to finish executing before it returns. This can be used to avoid waiting for the timeout to expire when
// gracefully shutting down your application. Calling Run at any point after calling Shutdown will lead to
// undefined behaviour.
func (b *Batcher) Shutdown(wait bool) { _ = "STUB: not implemented"; return }

func (b *Batcher) flushCurrentBatch() { _ = "STUB: not implemented"; return }

// stop the timer to avoid spurious flushes and trigger immediate cleanup in case this flush was
// triggered manually by a call to Shutdown (it has to happen inside the lock, so it can't be done
// in the Shutdown method directly)
