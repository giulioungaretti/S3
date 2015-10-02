//+build ignore
// Simple example that will nuke the contents of a s3 bucket
package main

import (
	"fmt"
	"runtime"
	"s3"
)

func init() {
	// multicorebitches
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func printKey(kb s3.KeyBucket) {
	fmt.Println(kb.Key.Key)
}

func main() {
	c := s3.Connection{}
	err := c.Connect()
	if err != nil {
		panic(err)
	}
	b, err := s3.Getbucket("eta-events-msgpack", &c)
	if err != nil {
		panic(err)
	}
	s3.ApplyToMultiList(b, "", "", printKey)
}
