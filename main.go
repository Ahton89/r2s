package main

import "github.com/Ahton89/r2s/pkg/r2s"

func main() {
	s := r2b.New()
	go s.RunRedisWorkers()
	r2b.Init(s)
}
