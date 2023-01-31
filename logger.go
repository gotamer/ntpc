package main

import (
	"log"
	"log/syslog"
	"os"
)

const APPNAME = "NPTC "

var (
	Debug  = *log.Default()
	Info  = *log.Default()
	Warn  = *log.Default()
	Error = *log.Default()
)

func init() {
	Debug.SetFlags(log.Lshortfile)
	Info.SetFlags(log.Lshortfile)
	Warn.SetFlags(log.Lshortfile)
	Error.SetFlags(log.Lshortfile)
}

func logger() {
	if os.Getenv("SHELL") == "/bin/sh" {
		Debug.Println("System Logger On")
		syslogger()
	}else{
		Debug.SetPrefix("DEBUG ")
		Info.SetPrefix("INFO ")
		Warn.SetPrefix("WARN ")
		Error.SetPrefix("ERROR ")
		Debug.Println("System Logger Off")
	}
}

func syslogger() {

	sysloggerD, err := syslog.New(syslog.LOG_CRON|syslog.LOG_DEBUG, APPNAME)
	if err != nil {
		log.Fatalln(err)
	}
	Debug.SetOutput(sysloggerD)

	sysloggerI, err := syslog.New(syslog.LOG_CRON|syslog.LOG_INFO, APPNAME)
	if err != nil {
		log.Fatalln(err)
	}
	Info.SetOutput(sysloggerI)

	sysloggerW, err := syslog.New(syslog.LOG_CRON|syslog.LOG_WARNING, APPNAME)
	if err != nil {
		log.Fatalln(err)
	}
	Warn.SetOutput(sysloggerW)

	sysloggerE, err := syslog.New(syslog.LOG_CRON|syslog.LOG_ERR, APPNAME)
	if err != nil {
		log.Fatalln(err)
	}
	Error.SetOutput(sysloggerE)
}
