package stats

import "sync"

var stats Stats

// thread-safe stats
type Stats struct {
	numMessagesSent int
	mu              sync.Mutex
}

func IncrementMessagesSent() {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.numMessagesSent++
}

func GetNumMessagesSent() int {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	return stats.numMessagesSent
}
