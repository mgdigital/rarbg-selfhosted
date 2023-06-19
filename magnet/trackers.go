package magnet

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetTrackers(path string) ([]string, error) {
	log.Info("loading trackers file")
	file, err := os.Open(path)
	if err != nil {
		log.WithField("error", err).Error("failed to open trackers file")
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.WithField("error", err).Error("failed to close trackers file")
			panic(err)
		}
	}(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.WithField("error", err).Error("failed to read from trackers file")
		return nil, err
	}
	return lines, err
}
