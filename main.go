package main

import (
	"log"
	"os"

	"github.com/zollidan/tornet/torrentfile"
)



func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	
	log.SetOutput(logFile)
	
	defer logFile.Close()

	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal("Error opening torrent file:", err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal("Error downloading torrent file:", err)
	}
}

