package main

import "time"

type TestResult struct {
	// CpuTime       time.Duration	// golang cannot get CPU time
	MemoryUsed    int64
	ExecutionTime time.Duration
	FinishedTime  time.Time
}
