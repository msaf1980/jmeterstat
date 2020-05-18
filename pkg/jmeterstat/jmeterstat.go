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
func (s *JMeterStat) Init() *JMeterStat {
	s.Started = 0
	s.Ended = 0

	s.Elapsed.Init()
	s.Connect.Init()
	s.Bytes.Init()
	s.SentBytes.Init()

	s.Success = 0
	s.ResponceCodes = map[string]uint64{}

	return s
}

// Add stat record to JMeterStat
func (s *JMeterStat) Add(timeStamp int64, elapsed float64, connect float64,
	bytes float64, sentBytes float64,
	success bool, responceCode string) *JMeterStat {

	if s.Started > timeStamp || s.Started == 0 {
		s.Started = timeStamp
	}
	if s.Ended < timeStamp {
		s.Ended = timeStamp
	}

	s.Elapsed.AddValue(elapsed)
	s.Connect.AddValue(connect)

	s.Bytes.AddValue(bytes)
	s.SentBytes.AddValue(sentBytes)

	if success {
		s.Success++
	}
	s.ResponceCodes[responceCode]++

	return s
}

// Equal with other instance
func (s *JMeterStat) Equal(o *JMeterStat) bool {
	if s.Started != o.Started {
		return false
	} else if s.Ended != o.Ended {
		return false
	} else if s.Success != o.Success {
		return false
	} else if !s.Elapsed.Equal(&o.Elapsed) {
		return false
	} else if !s.Connect.Equal(&o.Connect) {
		return false
	} else if !s.Bytes.Equal(&o.Bytes) {
		return false
	} else if !s.SentBytes.Equal(&o.SentBytes) {
		return false
	} else if len(s.ResponceCodes) != len(o.ResponceCodes) {
		return false
	} else {
		if len(s.ResponceCodes) == 0 {
			return true
		}
		for k, v := range s.ResponceCodes {
			vo, ok := o.ResponceCodes[k]
			if !ok || v != vo {
				return false
			}
		}
	}
	return true
}
