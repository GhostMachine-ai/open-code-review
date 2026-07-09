package stdout

import (
	"io"
	"os"
	"sync"
)

var (
	w  io.Writer = os.Stdout
	mu sync.RWMutex
)

// Writer returns the current stdout writer (real stdout or discard).
func Writer() io.Writer {
	mu.RLock()
	defer mu.RUnlock()
	return w
}

// Quiet replaces stdout with io.Discard and returns a cleanup function.
// Usage:
//
//	defer stdout.Quiet()()
//
// WARNING: Quiet must ONLY be called from the main goroutine, before spawning
// any concurrent work that writes to stdout, and its returned cleanup must be
// deferred in the same goroutine. Never call Quiet from multiple goroutines
// concurrently — it is not designed for nested or parallel silencing.
func Quiet() func() {
	return Redirect(io.Discard)
}

// Redirect swaps the progress writer for target and returns a cleanup
// function that restores the previous one. Callers use this to divert
// progress logs (e.g. to os.Stderr) instead of discarding them, so a
// consumer that parses JSON on stdout can still surface progress in
// real time.
//
// The same threading rules as Quiet apply: call only from the main
// goroutine before spawning writers, and defer its returned restore in
// the same goroutine.
func Redirect(target io.Writer) func() {
	mu.Lock()
	old := w
	w = target
	mu.Unlock()
	return func() {
		mu.Lock()
		w = old
		mu.Unlock()
	}
}
