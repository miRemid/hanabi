package config

import (
	"io"
	"log"
	"os"
)

func createdir(dir string) (bool, error) {
	_, err := os.Stat(dir)

	if err == nil {
		//directory exists
		return true, nil
	}

	err2 := os.MkdirAll(dir, 0755)
	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func setlog() {
	r, err := createdir(LOG_PATH)
	if r == false {
		panic(err)
	}
	file, _ := os.OpenFile(LOG_PATH+"/hanabi.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	log.SetOutput(io.MultiWriter(writers...))
	log.SetPrefix("[Hana]")
	log.SetFlags(log.Ldate | log.Ltime)
}
