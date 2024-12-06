package screamlogger

import "log"

func LoudScreamInfo(message string) {
	log.Printf("AAAAA INFO - %v", message)
}

func QuietScreamInfo(message string) {
	log.Printf("aaaaa INFO - %v", message)
}
