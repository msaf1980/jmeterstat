package jmeterreader

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
