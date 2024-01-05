package main

import (
	"crypto/md5"
	"crypto/sha256"
	"hash"
	"math/rand"
)

const dataLength = 10000000
const randomSeed = 28

type TestSha256 struct {
	TestCase
	data   []byte
	sha256 hash.Hash
}

func NewTestSha256() *TestSha256 {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	generator.Read(data)

	return &TestSha256{
		TestCase: NewTestCase(),
		data:     data,
		sha256:   sha256.New(),
	}
}

func (test *TestSha256) Run(loopCount int) {
	test.StartBenchmarking()
	test.runTest(loopCount)
	test.StopBenchmarking()
}

func (test *TestSha256) runTest(loopCount int) {
	for i := 0; i < loopCount; i++ {
		test.sha256.Sum(test.data)
	}
}

type TestMd5 struct {
	TestCase
	data []byte
	md5  hash.Hash
}

func NewTestMd5() *TestMd5 {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	generator.Read(data)

	return &TestMd5{
		TestCase: NewTestCase(),
		data:     data,
		md5:      md5.New(),
	}
}

func (test *TestMd5) Run(loopCount int) {
	test.StartBenchmarking()
	test.runTest(loopCount)
	test.StopBenchmarking()
}

func (test *TestMd5) runTest(loopCount int) {
	for i := 0; i < loopCount; i++ {
		test.md5.Sum(test.data)
	}
}
