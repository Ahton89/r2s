package r2b

import (
	"github.com/Sirupsen/logrus"
	"os"
)

func (s *R2s) Migrate(){
	for _,hash := range s.config.HashesToCopy {
		s.keysCount = 0
		hashKeys,err := s.getHashKeys(hash)
		if err != nil {
			s.log.WithFields(logrus.Fields{"hash":hash}).Error("Can not get keys list for hash")
			os.Exit(1)
		}
		for _,key := range hashKeys {
			s.workerChannel <- &copyStruct{hash, key}
			s.total++
			s.keysCount++
			s.wg.Add(1)
		}
		if !*s.debug {
			s.log.WithFields(logrus.Fields{"hash":hash, "keys":s.keysCount}).Info("Cloning...")
		}
	}
	s.wg.Wait()
}

func (s *R2s) RunRedisWorkers() {
	for i := 0; i < s.workersCount; i++ {
		go s.runWorker()
	}
}


func (s *R2s) runWorker() {
	for elem := range s.workerChannel {
		s.Clone(elem)
	}
}

func (s *R2s) Clone(a *copyStruct) {
	setHashError := s.setHash(a.Hash, a.Key, s.getHashValues(a.Hash, a.Key))
	s.log.WithFields(logrus.Fields{"hash":a.Hash, "key":a.Key}).Debug("Cloning...")
	if setHashError != nil {
		s.log.WithFields(logrus.Fields{"hash":a.Hash, "key":a.Key, "error":setHashError}).Error("Can not clone :(")
	}
	s.wg.Done()
}
