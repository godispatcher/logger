package logger

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	logfile *os.File
	mutex   sync.Mutex
)

func InitLogFile(path string) error {
	var err error
	logfile, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}
func CloseLogFile() {
	logfile.Close()
}
func WriteLog(entry LogEntry) error {
	mutex.Lock()
	defer mutex.Unlock()

	entryJSON, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = logfile.WriteString(string(entryJSON) + "\n")
	return err
}
