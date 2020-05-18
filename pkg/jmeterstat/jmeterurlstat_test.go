package jmeterstat

import (
	"testing"

	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
	"github.com/msaf1980/jmeterstat/pkg/statcalc"
)

func jMeterURLStatEqual(s JMeterURLStat, o JMeterURLStat) bool {
	if len(s) != len(o) {
		return false
	} else if len(s) == 0 {
		return true
	}
	for k, v := range s {
		vo, ok := o[k]
		if !ok {
			return false
		}
		if !v.Equal(vo) {
			return false
		}
	}
	return true
}

func jMeterLabelURLStatEqual(s JMeterLabelURLStat, o JMeterLabelURLStat) bool {
	if len(s) != len(o) {
		return false
	} else if len(s) == 0 {
		return true
	}
	for l, u := range s {
		uo, ok := o[l]
		if !ok {
			return false
		}
		if !jMeterURLStatEqual(u, uo) {
			return false
		}
	}
	return true
}

func TestJMeterURLStatAdd(t *testing.T) {
	urlStat := JMeterLabelURLStat{}

	tests := []struct {
		name   string
		input  jmeterreader.JmtrRecord
		result JMeterLabelURLStat
	}{
		{
			"1. Add /test1, code 200",
			jmeterreader.JmtrRecord{
				TimeStamp:    1,
				Elapsed:      2.0,
				Label:        "test1",
				ResponseCode: "200",
				Success:      true,
				Bytes:        10,
				SentBytes:    120,
				GrpThreads:   1,
				AllThreads:   2,
				URL:          "/test1",
				Latency:      2.0,
				IdleTime:     0.0,
				Connect:      1.0,
			},
			JMeterLabelURLStat{
				"test1": {
					"/test1": new(JMeterStat).Init().Add(1, 2.0, 1.0, 10.0, 120.0, true, "200"),
				},
			},
		},
		{
			"2. Add /test1, code 500",
			jmeterreader.JmtrRecord{
				TimeStamp:    2,
				Elapsed:      2.0,
				Label:        "test1",
				ResponseCode: "500",
				Success:      false,
				Bytes:        10,
				SentBytes:    120,
				GrpThreads:   1,
				AllThreads:   2,
				URL:          "/test1",
				Latency:      2.0,
				IdleTime:     0.0,
				Connect:      1.0,
			},
			JMeterLabelURLStat{
				"test1": {
					"/test1": new(JMeterStat).Init().Add(1, 2.0, 1.0, 10.0, 120.0, true, "200").
						Add(2, 2.0, 1.0, 10.0, 120.0, false, "500"),
				},
			},
		},
		{
			"3. Add /test2, code 200",
			jmeterreader.JmtrRecord{
				TimeStamp:    3,
				Elapsed:      3.0,
				Label:        "test2",
				ResponseCode: "200",
				Success:      true,
				Bytes:        30,
				SentBytes:    60,
				GrpThreads:   1,
				AllThreads:   2,
				URL:          "/test2",
				Latency:      3.0,
				IdleTime:     0.0,
				Connect:      1.0,
			},
			JMeterLabelURLStat{
				"test1": {
					"/test1": new(JMeterStat).Init().Add(1, 2.0, 1.0, 10.0, 120.0, true, "200").
						Add(2, 2.0, 1.0, 10.0, 120.0, false, "500"),
				},
				"test2": {
					"/test2": new(JMeterStat).Init().Add(3, 3.0, 1.0, 30.0, 60.0, true, "200"),
				},
			},
		},
	}
	for _, tt := range tests {
		JMeterURLStatAdd(urlStat, tt.input.URL, &tt.input)
		if !jMeterLabelURLStatEqual(urlStat, tt.result) {
			t.Errorf("%s failed", tt.name)
		}
	}
}

func TestJMeterURLStatSum(t *testing.T) {
	tests := []struct {
		name   string
		input  JMeterURLStat
		result JMeterStat
	}{
		{
			"1. Summ /test1",
			JMeterURLStat{
				"/test1": new(JMeterStat).Init().Add(1, 2.0, 1.0, 10.0, 120.0, true, "200").
					Add(2, 3.0, 1.0, 9.0, 100.0, false, "500"),
			},
			JMeterStat{
				1, 2,
				*(new(statcalc.StatCalculator).Init().AddValue(2.0).AddValue(3.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0).AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(10.0).AddValue(9.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(120.0).AddValue(100.0)),
				1,
				map[string]uint64{"200": 1, "500": 1},
			},
		},
	}
	for _, tt := range tests {
		got := JMeterURLStatSum(tt.input, false)
		if !got.Equal(&tt.result) {
			t.Errorf("%s failed, want %v, got %v", tt.name, tt.result, got)
		}
	}
}
