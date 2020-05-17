package jmetercsv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// JMeter CSV fileds index numbers
type JmtrCsvHeader struct {
	Length          int // fields count
	TimeStamp       int8
	Elapsed         int8
	Label           int8
	ResponseCode    int8
	ResponseMessage int8
	ThreadName      int8
	DataType        int8
	Success         int8
	FailureMessage  int8
	Bytes           int8
	SentBytes       int8
	GrpThreads      int8
	AllThreads      int8
	URL             int8
	Latency         int8
	IdleTime        int8
	Connect         int8
}

// Set CSV fields index numbers
func jmtrCsvGetHeader(line []string, head *JmtrCsvHeader) (err error) {
	head.TimeStamp = -1
	head.Elapsed = -1
	head.Label = -1
	head.ResponseCode = -1
	head.ResponseMessage = -1
	head.ThreadName = -1
	head.DataType = -1
	head.Success = -1
	head.FailureMessage = -1
	head.Bytes = -1
	head.SentBytes = -1
	head.GrpThreads = -1
	head.AllThreads = -1
	head.URL = -1
	head.Latency = -1
	head.IdleTime = -1
	head.Connect = -1
	head.Length = 17

	for i := range line {
		switch line[i] {
		case "timeStamp":
			head.TimeStamp = int8(i)
		case "elapsed":
			head.Elapsed = int8(i)
		case "label":
			head.Label = int8(i)
		case "responseCode":
			head.ResponseCode = int8(i)
		case "responseMessage":
			head.ResponseMessage = int8(i)
		case "threadName":
			head.ThreadName = int8(i)
		case "dataType":
			head.DataType = int8(i)
		case "success":
			head.Success = int8(i)
		case "failureMessage":
			head.FailureMessage = int8(i)
		case "bytes":
			head.Bytes = int8(i)
		case "sentBytes":
			head.SentBytes = int8(i)
		case "grpThreads":
			head.GrpThreads = int8(i)
		case "allThreads":
			head.AllThreads = int8(i)
		case "URL":
			head.URL = int8(i)
		case "Latency":
			head.Latency = int8(i)
		case "IdleTime":
			head.IdleTime = int8(i)
		case "Connect":
			head.Connect = int8(i)
		default:
			err = fmt.Errorf("unknown field: %s", line[i])
			return
		}
	}
	if head.TimeStamp == -1 {
		err = fmt.Errorf("missing field: timeStamp")
		return
	}
	if head.Elapsed == -1 {
		err = fmt.Errorf("missing field: elapsed")
		return
	}
	if head.Label == -1 {
		err = fmt.Errorf("missing field: label")
		return
	}
	if head.ResponseCode == -1 {
		err = fmt.Errorf("missing field: timeStamp")
		return
	}
	if head.ResponseMessage == -1 {
		err = fmt.Errorf("missing field: responseMessage")
		return
	}
	if head.ThreadName == -1 {
		err = fmt.Errorf("missing field: threadName")
		return
	}
	if head.DataType == -1 {
		err = fmt.Errorf("missing field: dataType")
		return
	}
	if head.Success == -1 {
		err = fmt.Errorf("missing field: success")
		return
	}
	if head.FailureMessage == -1 {
		err = fmt.Errorf("missing field: failureMessage")
		return
	}
	if head.Bytes == -1 {
		err = fmt.Errorf("missing field: bytes")
		return
	}
	if head.SentBytes == -1 {
		err = fmt.Errorf("missing field: sentBytes")
		return
	}
	if head.GrpThreads == -1 {
		err = fmt.Errorf("missing field: grpThreads")
		return
	}
	if head.AllThreads == -1 {
		err = fmt.Errorf("missing field: allThreads")
		return
	}
	if head.URL == -1 {
		err = fmt.Errorf("missing field: URL")
		return
	}
	if head.Latency == -1 {
		err = fmt.Errorf("missing field: Latency")
		return
	}
	if head.IdleTime == -1 {
		err = fmt.Errorf("missing field: IdleTime")
		return
	}
	if head.Connect == -1 {
		err = fmt.Errorf("missing field: Connect")
		return
	}
	if len(line) != head.Length {
		err = fmt.Errorf("mismatch fields count in header")
		return
	}
	return
}

// JMter CSV Reader
type JmtrCsvReader struct {
	Header    JmtrCsvHeader
	csvFile   *os.File
	csvReader *csv.Reader
}

// Init Jmeter CSV Reader (aloocate and read header)
func NewJmtrCsvReader(csvFilename *string) (*JmtrCsvReader, error) {
	var err error
	p := new(JmtrCsvReader)
	p.csvFile, err = os.Open(*csvFilename)
	if err != nil {
		return nil, err
	}
	p.csvReader = csv.NewReader(bufio.NewReader(p.csvFile))
	p.csvReader.ReuseRecord = true
	line, err := p.csvReader.Read()
	if err != nil {
		p.csvFile.Close()
		return nil, err
	}
	err = jmtrCsvGetHeader(line, &p.Header)
	if err != nil {
		p.csvFile.Close()
		return nil, err
	}
	return p, nil
}

// JMeter stat record
type JmtrRecord struct {
	TimeStamp    int64
	Elapsed      float64
	Label        string
	ResponseCode string
	//ResponseMessage string
	//ThreadName string
	//DataType   string
	Success bool
	//FailureMessage  string
	Bytes      int64
	SentBytes  int64
	GrpThreads int64
	AllThreads int64
	URL        string
	Latency    float64
	IdleTime   float64
	Connect    float64
}

// Read line from Jmeter CSV file
func (p *JmtrCsvReader) Read(r *JmtrRecord) error {
	line, err := p.csvReader.Read()
	if err != nil {
		p.csvFile.Close()
		return err
	}
	r.TimeStamp, err = strconv.ParseInt(line[p.Header.TimeStamp], 10, 64)
	if err != nil || r.TimeStamp <= 0 {
		err = fmt.Errorf("incorrect timeStamp value: %s", line[p.Header.TimeStamp])
		return err
	}
	r.Elapsed, err = strconv.ParseFloat(line[p.Header.Elapsed], 64)
	if err != nil || r.Elapsed < 0 {
		err = fmt.Errorf("incorrect elapsed value: %s", line[p.Header.Elapsed])
		return err
	}
	r.Label = line[p.Header.Label]
	r.ResponseCode = line[p.Header.ResponseCode]
	//ResponseMessage
	//ThreadName
	//DataType
	if line[p.Header.Success] == "true" {
		r.Success = true
	} else if line[p.Header.Success] == "false" {
		r.Success = false
	} else {
		err = fmt.Errorf("incorrect success value: %s", line[p.Header.Success])
		return err
	}
	//FailureMessage
	r.Bytes, err = strconv.ParseInt(line[p.Header.Bytes], 10, 64)
	if err != nil || r.Bytes < 0 {
		err = fmt.Errorf("incorrect bytes value: %s", line[p.Header.Bytes])
		return err
	}
	r.SentBytes, err = strconv.ParseInt(line[p.Header.SentBytes], 10, 64)
	if err != nil || r.SentBytes < 0 {
		err = fmt.Errorf("incorrect sentBytes value: %s", line[p.Header.SentBytes])
		return err
	}
	r.GrpThreads, err = strconv.ParseInt(line[p.Header.GrpThreads], 10, 64)
	if err != nil || r.GrpThreads < 0 {
		err = fmt.Errorf("incorrect grpThreads value: %s", line[p.Header.GrpThreads])
		return err
	}
	r.AllThreads, err = strconv.ParseInt(line[p.Header.AllThreads], 10, 64)
	if err != nil || r.AllThreads < 0 {
		err = fmt.Errorf("incorrect allThreads value: %s", line[p.Header.AllThreads])
		return err
	}
	//AllThreads      int
	r.URL = line[p.Header.URL]
	r.Latency, err = strconv.ParseFloat(line[p.Header.Latency], 64)
	if err != nil || r.Latency < 0 {
		err = fmt.Errorf("incorrect Latency value: %s", line[p.Header.Latency])
		return err
	}
	r.IdleTime, err = strconv.ParseFloat(line[p.Header.IdleTime], 64)
	if err != nil || r.GrpThreads < 0 {
		err = fmt.Errorf("incorrect IdleTime value: %s", line[p.Header.IdleTime])
		return err
	}
	r.Connect, err = strconv.ParseFloat(line[p.Header.Connect], 64)
	if err != nil || r.GrpThreads < 0 {
		err = fmt.Errorf("incorrect GrpThreads value: %s", line[p.Header.Connect])
		return err
	}
	return nil
}
