package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// Create a .log file in the specified path and log the file name and line number where the log is created
func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum", LstdFlags)
	iLog.Println("Learning Go paps")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another cool log entry paps")
}
