package jmeterstat

import (
	"testing"

	"github.com/msaf1980/jmeterstat/pkg/statcalc"
)

func TestJMeterStat(t *testing.T) {
	type args struct {
		timeStamp    int64
		elapsed      float64
		connect      float64
		bytes        float64
		sentBytes    float64
		success      bool
		responceCode string
	}
	tests := []struct {
		name   string
		reset  bool
		args   args
		result JMeterStat
		equal  bool
	}{
		{
			"1. JMeterStat.Add(10, 10.0, 1.0, 11.0, 120.0, true, '200')",
			false,
			args{10, 10.0, 1.0, 11.0, 120.0, true, "200"},
			JMeterStat{
				10, 10,
				*(new(statcalc.StatCalculator).Init().AddValue(10.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(11.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(120.0)),
				1,
				map[string]uint64{"200": 1},
				map[string]uint64{},
			},
			true,
		},
		{
			"2. JMeterStat.Add(11, 10.0, 1.0, 11.0, 120.0, true, '200')",
			false,
			args{11, 10.0, 1.0, 11.0, 120.0, true, "200"},
			JMeterStat{
				10, 11,
				*(new(statcalc.StatCalculator).Init().AddValue(10.0).AddValue(10.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0).AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(11.0).AddValue(11.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(120.0).AddValue(120.0)),
				2,
				map[string]uint64{"200": 2},
				map[string]uint64{},
			},
			true,
		},
		{
			"3. JMeterStat.Add(12, 10.0, 1.0, 11.0, 120.0, false, '500')",
			false,
			args{12, 10.0, 1.0, 11.0, 120.0, false, "500"},
			JMeterStat{
				10, 12,
				*(new(statcalc.StatCalculator).Init().AddValue(10.0).AddValue(10.0).AddValue(10.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0).AddValue(1.0).AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(11.0).AddValue(11.0).AddValue(11.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(120.0).AddValue(120.0).AddValue(120.0)),
				2,
				map[string]uint64{"200": 2},
				map[string]uint64{"500": 1},
			},
			true,
		},
		{
			"JMeterStat.Init().Add(10, 10, 1, 5.0, 12.0, false, '500')",
			true,
			args{10, 10.0, 1.0, 5.0, 12.0, false, "500"},
			JMeterStat{
				10, 10,
				*(new(statcalc.StatCalculator).Init().AddValue(10.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(11.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(120.0)),
				0,
				map[string]uint64{},
				map[string]uint64{"500": 1},
			},
			false,
		},
		{
			"JMeterStat.Init().Add(12, 10, 1, 5.0, 12.0, false, '500')",
			true,
			args{12, 10.0, 1.0, 5.0, 12.0, false, "500"},
			JMeterStat{
				12, 12,
				*(new(statcalc.StatCalculator).Init().AddValue(10.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(1.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(5.0)),
				*(new(statcalc.StatCalculator).Init().AddValue(12.0)),
				0,
				map[string]uint64{},
				map[string]uint64{"500": 1},
			},
			true,
		},
	}

	var stat JMeterStat
	stat.Init()

	for _, tt := range tests {
		if tt.reset {
			stat.Init()
		}
		stat.Add(tt.args.timeStamp, tt.args.elapsed, tt.args.connect, tt.args.bytes, tt.args.sentBytes,
			tt.args.success, tt.args.responceCode)
		got := stat.Equal(&tt.result)
		if got != tt.equal {
			t.Errorf("%s = %v, want %v", tt.name, got, tt.equal)
		}
	}
}
