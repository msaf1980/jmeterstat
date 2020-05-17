package jmeterstat

import (
	"github.com/msaf1980/jmeterstat/pkg/statcalc"
)

// JMeterStat aggegator for jmeter stat
type JMeterStat struct {
	Started int64
	Ended   int64

	Elapsed   statcalc.StatCalculator
	Connect   statcalc.StatCalculator
	Bytes     statcalc.StatCalculator
	SentBytes statcalc.StatCalculator

	Success       uint64
	ResponceCodes map[string]uint64
}

// Init reset JMeterStat to initial empthy state. Also call for next reuse.
func (s *JMeterStat) Init(started int64) {
	s.Started = started
	s.Ended = 0

	s.Elapsed.Init()
	s.Connect.Init()
	s.Bytes.Init()
	s.SentBytes.Init()
}

// Add stat record to JMeterStat
func (s *JMeterStat) Add(elapsed float64, connect float64, bytes float64, sentBytes float64) {
	s.Elapsed.AddValue(elapsed)
	s.Connect.AddValue(connect)
	s.Bytes.AddValue(bytes)
	s.SentBytes.AddValue(sentBytes)
}
