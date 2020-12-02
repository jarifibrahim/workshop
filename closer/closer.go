package main

import (
	"context"
	"sync"
)

// Taken from https://github.com/dgraph-io/ristretto/blob/67fef616c676b6848c3fd026d16b8f7d7ef6ae87/z/z.go#L65-L134

// Closer holds the two things we need to close a goroutine and wait for it to
// finish: a chan to tell the goroutine to shut down, and a WaitGroup with
// which to wait for it to finish shutting down.
//
// How to use Closer?
// -> Create a new closer.
// -> Pass the closer to the goroutine.
// -> Inside the goroutine, check if the closer has been closed by HasBeenClosed() function.
// 		-> If closed, mark closer as done using closer.Done() and return from the goroutine.
// -> To stop the goroutine, use the closer.Signal() method.
// -> To stop and wait for the goroutine to finish, call closer.SignalAndWait() method.
type Closer struct {
	waiting sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc
}

// NewCloser constructs a new Closer, with an initial count on the WaitGroup.
func NewCloser(initial int) *Closer {
	ret := &Closer{}
	ret.ctx, ret.cancel = context.WithCancel(context.Background())
	ret.waiting.Add(initial)
	return ret
}

// AddRunning Add()'s delta to the WaitGroup.
func (lc *Closer) AddRunning(delta int) {
	lc.waiting.Add(delta)
}

// Ctx can be used to get a context, which would automatically get cancelled when Signal is called.
func (lc *Closer) Ctx() context.Context {
	return lc.ctx
}

// Signal signals the HasBeenClosed signal.
func (lc *Closer) Signal() {
	lc.cancel()
}

// HasBeenClosed gets signaled when Signal() is called.
func (lc *Closer) HasBeenClosed() <-chan struct{} {
	return lc.ctx.Done()
}

// Done calls Done() on the WaitGroup.
func (lc *Closer) Done() {
	lc.waiting.Done()
}

// Wait waits on the WaitGroup. (It waits for NewCloser's initial value, AddRunning, and Done
// calls to balance out.)
func (lc *Closer) Wait() {
	lc.waiting.Wait()
}

// SignalAndWait calls Signal(), then Wait().
func (lc *Closer) SignalAndWait() {
	lc.Signal()
	lc.Wait()
}
