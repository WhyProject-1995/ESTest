package main

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Panic(log)
	}

	hook, err := elogrus.NewElasticHook(client, "localhost", log.Level, "mylog")
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(hook)
}

func main() {
	for i := 0; i < 100; i++ {
		log.Debugf("debug: %d", i)
	}
	log.WithFields(logrus.Fields{}).Info("hello world")
	log.WithFields(logrus.Fields{}).Info("A group of walrus emerges from the ocean")
}
