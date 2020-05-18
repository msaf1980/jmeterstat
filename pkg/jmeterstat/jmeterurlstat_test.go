package jmeterstat

import (
	"reflect"
	"testing"

	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
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
				1,   //TimeStamp
				2.0, //Elapsed
				"test1",
				"200", true,
				10, 120,
				1,
				2,
				"/test1",
				2.0, 0.0, 1.0,
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
				2,   //TimeStamp
				2.0, //Elapsed
				"test1",
				"500", false,
				10, 120,
				1,
				2,
				"/test1",
				2.0, 0.0, 1.0,
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
				3, 3.0, // Elapsed
				"test2",
				"200", true,
				30, 60,
				1,
				2,
				"/test2",
				3.0, 0.0, 1.0,
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
			t.Errorf("%s failed, want %v, got %v", tt.name, tt.result, urlStat)
		}
	}
}

func TestJMeterURLStatSum(t *testing.T) {
	type args struct {
		stat  JMeterURLStat
		clear bool
	}
	tests := []struct {
		name string
		args args
		want *JMeterStat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JMeterURLStatSum(tt.args.stat, tt.args.clear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JMeterURLStatSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
