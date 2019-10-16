package main

import "github.com/Ahton89/r2s/pkg/r2s"

var Version = ""

func main() {
	s := r2b.New(Version)
	go s.RunRedisWorkers()
	r2b.Init(s)
}
