package main

import "time"

type TestResult struct {
	CpuTime       time.Duration
	MemoryUsed    uint64
	ExecutionTime time.Duration
	FinishedTime  time.Time
}
