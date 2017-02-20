package main

import (
	"flag"
	"os"

	"strings"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	logLevel := flag.String("log-level", "debug", "Log level")
	logFile := flag.String("log-file", "", "Log file")
	flag.Parse()

	*logLevel = strings.ToLower(*logLevel)

	switch *logLevel {
	default:
		log.Fatal("Bad input for 'log-level' flag")
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "warn", "warning":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "panic":
		log.Level = logrus.PanicLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	}
	// logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	if *logFile != "" {
		file, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.WithFields(logrus.Fields{
				"file": *logFile,
			}).Info("Failed to log to file, using default stderr")
		}
	} else {
		log.Info("No logfile provided, using default stderr")
	}
	log.Infof("Logging on %v level", log.Level)
}
