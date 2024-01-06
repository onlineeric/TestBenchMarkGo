package main

import (
	"crypto/md5"
	"crypto/sha256"
	"math/rand"
)

const dataLength = 10000000
const randomSeed = 28

type TestSha256 struct {
	TestCase
	data []byte
}

func NewTestSha256() *TestSha256 {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	generator.Read(data)

	return &TestSha256{
		TestCase: NewTestCase(),
		data:     data,
	}
}

func (test *TestSha256) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		sha256.Sum256(test.data)
	}
	test.StopBenchmarking()
}

type TestMd5 struct {
	TestCase
	data []byte
}

func NewTestMd5() *TestMd5 {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	generator.Read(data)

	return &TestMd5{
		TestCase: NewTestCase(),
		data:     data,
	}
}

func (test *TestMd5) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		md5.Sum(test.data)
	}
	test.StopBenchmarking()
}
