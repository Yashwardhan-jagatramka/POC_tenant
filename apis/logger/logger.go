package logger

import "os"

func CreateLogFile() *os.File {
	var LogFile, _ = os.Create("logfile.txt")
	return LogFile
}

func CloseLogFile() {
	defer CreateLogFile().Close()
}
