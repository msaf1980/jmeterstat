package main

import (
	"flag"
	"time"
)

type JmtrCsvHeader struct {
	timeStamp       int
	elapsed         int
	label           int
	responseCode    int
	responseMessage int
	threadName      int
	dataType        int
	success         int
	failureMessage  int
	bytes           int
	sentBytes       int
	grpThreads      int
	allThreads      int
	URL             int
	Latency         int
	IdleTime        int
	Connect         int
}

type JmtrCsvRecod struct {
	timeStamp    time.Time
	elapsed      int
	label        string
	responseCode string
	//responseMessage string
	threadName string
	dataType   string
	success    bool
	//failureMessage  string
	bytes      int
	sentBytes  int
	grpThreads int
	//allThreads      int
	URL      string
	Latency  int
	IdleTime int
	Connect  int
}
