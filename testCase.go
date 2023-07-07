package main

import (
	"runtime"
	"time"
)

type TestCase struct {
	startCpuTime   time.Duration
	startMemory    uint64
	Result         *TestResult
	stopwatchStart time.Time
}

func NewTestCase() TestCase {
	return TestCase{}
}

func (tc *TestCase) StartBenchmarking() {
	var memStats runtime.MemStats

	tc.startCpuTime = time.Now().Sub(time.Now())
	runtime.ReadMemStats(&memStats)
	tc.startMemory = memStats.Alloc

	tc.stopwatchStart = time.Now()
}

func (tc *TestCase) StopBenchmarking() {
	var memStats runtime.MemStats

	runtime.ReadMemStats(&memStats)
	endMemory := memStats.Alloc
	endCpuTime := time.Since(tc.stopwatchStart)

	tc.Result = &TestResult{
		CpuTime:       endCpuTime,
		MemoryUsed:    endMemory - tc.startMemory,
		ExecutionTime: endCpuTime,
		FinishedTime:  time.Now(),
	}
}
