package r2b

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/akamensky/argparse"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

const version = "0.0.3"

func New() *R2s {
	s := &R2s{
		log:           logrus.New(),
		arg:           argparse.NewParser("redis_p2s", fmt.Sprintf("Redis Production to Sandbox cloning tool v.%s", version)),
		workerChannel: make(chan *copyStruct),
		workersCount:  7,
		configFile:    fmt.Sprintf("%s/config.yml", path.Join(path.Dir(os.Args[0]))),
	}
	return s
}

func Init(s *R2s) {
	s.prodHost = s.arg.String("p", "production", &argparse.Options{Required: true, Help: "Redis production node <host:port>"})
	s.sandboxHost = s.arg.String("s", "sandbox", &argparse.Options{Required: true, Help: "Redis sandbox node <host:port>"})
	s.prodDb = s.arg.Int("i", "production-db", &argparse.Options{Required: false, Help: "Redis production db <num>", Default: 0})
	s.sandboxDb = s.arg.Int("o", "sandbox-db", &argparse.Options{Required: false, Help: "Redis sandbox db <num>", Default: 0})
	s.debug = s.arg.Flag("d", "debug", &argparse.Options{Help: "Enable debug mode"})
	err := s.arg.Parse(os.Args)
	if *s.debug {
		s.log.Level = logrus.DebugLevel
	}
	if err != nil {
		fmt.Print(s.arg.Usage(err))
		os.Exit(1)
	}
	configError := s.getConfig()
	if configError != nil {
		s.log.WithFields(logrus.Fields{"config": s.configFile}).Error(configError)
		os.Exit(1)
	}
	s.redisProd = s.redisConnect(*s.prodHost, *s.prodDb)
	s.redisSandbox = s.redisConnect(*s.sandboxHost, *s.sandboxDb)
	defer s.redisProd.Close()
	defer s.redisSandbox.Close()
	defer close(s.workerChannel)
	s.wg = new(sync.WaitGroup)
	s.log.WithFields(logrus.Fields{"production-host": *s.prodHost, "production-db": *s.prodDb, "sandbox": *s.sandboxHost, "sandbox-db": *s.sandboxDb}).Info("Starting redis cloning... ")
	s.Migrate()
}

func (s *R2s) getConfig() error {
	configByte, err := ioutil.ReadFile(s.configFile)
	if err != nil {
		return errors.New("can not read config file")
	}
	err = yaml.Unmarshal(configByte, &s.config)
	if err != nil {
		return errors.New("can not get config data, check config")
	}
	return err
}
