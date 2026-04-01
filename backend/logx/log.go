package logx

import "log"

// Simple wrapper to provide consistent log prefixes. Replace with structured logger later.

func Infof(format string, v ...interface{}) {
	log.Printf("INFO: "+format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}

func Debugf(format string, v ...interface{}) {
	log.Printf("DEBUG: "+format, v...)
}
