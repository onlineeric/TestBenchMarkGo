package main

import (
	"runtime"
	"time"
)

type TestCase struct {
	// startCpuTime   time.Duration		// golang cannot get CPU time
	startMemory    uint64
	Result         *TestResult
	stopwatchStart time.Time
}

func NewTestCase() TestCase {
	return TestCase{}
}

func (tc *TestCase) StartBenchmarking() {
	var memStats runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&memStats)
	tc.startMemory = memStats.Alloc

	tc.stopwatchStart = time.Now()
}

func (tc *TestCase) StopBenchmarking() {
	var memStats runtime.MemStats

	runtime.ReadMemStats(&memStats)
	endMemory := memStats.Alloc
	tc.Result = &TestResult{
		MemoryUsed:    calcMemoryUsed(tc.startMemory, endMemory),
		ExecutionTime: time.Since(tc.stopwatchStart),
		FinishedTime:  time.Now(),
	}
}

func calcMemoryUsed(startMemory uint64, endMemory uint64) int64 {
	if endMemory < startMemory {
		return int64(startMemory-endMemory) * -1
	}
	return int64(endMemory - startMemory)
}
