package hdfs

import (
	"github.com/vladimirvivien/gowfs"
	log "stupid-ddbs/logutil"
	"sync"
)

type Manager struct {
	client *gowfs.FileSystem
}

var instance *Manager
var once sync.Once

func GetManagerInstance() *Manager {
	once.Do(func() {
		log.Info("hdfs manager starts to initialize.")
		defer log.Info("hdfs manager has been initialized.")
		config := gowfs.Configuration{Addr: "localhost:9870", User: "root"}
		client, err := gowfs.NewFileSystem(config)
		if err != nil {
			panic(err)
		}
		instance = &Manager{
			client,
		}
	})
	return instance
}

func (m* Manager) Close() {
	log.Info("hdfs manager close (no action)")
}




