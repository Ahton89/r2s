package r2b

import (
	"github.com/akamensky/argparse"
	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
	"sync"
)

type R2s struct {
	arg           *argparse.Parser
	log           *logrus.Logger
	prodHost      *string
	sandboxHost   *string
	prodDb        *int
	sandboxDb     *int
	redisProd     *redis.Client
	redisSandbox  *redis.Client
	workerChannel chan *copyStruct
	workersCount  int
	wg            *sync.WaitGroup
	keyCount      int
	config        *config
	configFile    string
	debug         *bool
	total         int
	keysCount     int
}

type config struct {
	HashesToCopy []string `yaml:"hashesToCopy"`
}

type copyStruct struct {
	Hash string
	Key  string
}
