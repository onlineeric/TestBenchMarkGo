package main

import (
	"fmt"
)

func main() {
	sha256Test := NewTestSha256()
	sha256Test.Run(500)

	result := sha256Test.Result

	if result != nil {
		fmt.Printf("SHA256 CPU Time: %f ms\n", result.CpuTime.Seconds()*1000)
		fmt.Printf("SHA256 Memory Used: %d bytes\n", result.MemoryUsed)
		fmt.Printf("SHA256 Execution Time: %f ms\n", result.ExecutionTime.Seconds()*1000)
	} else {
		fmt.Println("SHA256 failed")
	}

	md5Test := NewTestMd5()
	md5Test.Run(500)

	result = md5Test.Result

	if result != nil {
		fmt.Printf("MD5 CPU Time: %f ms\n", result.CpuTime.Seconds()*1000)
		fmt.Printf("MD5 Memory Used: %d bytes\n", result.MemoryUsed)
		fmt.Printf("MD5 Execution Time: %f ms\n", result.ExecutionTime.Seconds()*1000)
	} else {
		fmt.Println("MD5 failed")
	}
}