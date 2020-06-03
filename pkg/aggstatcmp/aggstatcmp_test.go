package aggstatcmp

import (
	"math"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/msaf1980/jmeterstat/pkg/aggstat"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func Test_diffPcnt(t *testing.T) {
	type args struct {
		o   float64
		s   float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"NaN, 0.0, 10.0", args{math.NaN(), 0.0, 10.0}, math.NaN()},
		{"0.0, NaN, 10.0", args{0.0, math.NaN(), 10.0}, math.NaN()},
		{"NaN, 10.0, NaN", args{math.NaN(), math.NaN(), 10.0}, math.NaN()},
		{"0.1, 1.0, 10.0", args{0.1, 1.0, 10.0}, 9.0},
		{"1.0, 0.1, 10.0", args{1.0, 0.1, 10.0}, -9.0},
	}
	for _, tt := range tests {
		got := diffPcnt(tt.args.s, tt.args.o, tt.args.max)
		mismatch := false
		if math.IsNaN(tt.want) {
			if !math.IsNaN(got) {
				mismatch = true
			}
		} else if got != tt.want {
			mismatch = true
		}
		if mismatch {
			t.Errorf("%s diffPcnt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_diffMapPcnt(t *testing.T) {
	type args struct {
		o         map[string]uint64
		oRequests uint64
		s         map[string]uint64
		sRequests uint64
		want      map[string]float64
	}
	tests := []args{
		{
			map[string]uint64{"500": 1, "503": 2, "504": 4},
			100,
			map[string]uint64{"403": 20, "500": 30, "504": 50},
			1000,
			map[string]float64{"403": 2.0, "500": 2.0, "503": -2.0, "504": 1.0},
		},
	}
	for _, tt := range tests {
		got := map[string]float64{}
		diffMapPcnt(&got, &tt.s, tt.sRequests, &tt.o, tt.oRequests)
		mismatch := false
		if len(got) != len(tt.want) {
			mismatch = true
		} else {
			for k := range tt.want {
				v, ok := got[k]
				if ok {
					if !almostEqual(v, tt.want[k]) {
						mismatch = true
						break
					}
				} else {
					mismatch = true
					break
				}
			}
		}
		if mismatch {
			t.Errorf("diffMapPcnt() =\n%#v\nwant\n%#v", got, tt.want)
		}
	}
}

func TestLabelURLAggDiffStat_Diff(t *testing.T) {
	type args struct {
		s aggstat.LabelURLAggStat
		o aggstat.LabelURLAggStat
	}
	tests := []struct {
		name   string
		args   args
		result LabelURLAggDiffStat
	}{
		{
			"Equal",
			args{
				s: aggstat.LabelURLAggStat{Started: 1, Ended: 2, Name: "Test",
					Stat: map[string]*aggstat.URLAggStat{
						"test1": {
							Stat: map[string]*aggstat.AggStat{
								"/test1": &aggstat.AggStat{
									Started: 1, Ended: 2, Count: 1,
									Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
									Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
									Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
									SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
									Success:      1,
									SuccessCodes: map[string]uint64{"200": 1},
									ErrorCodes:   map[string]uint64{"500": 1},
								},
								"/test2": &aggstat.AggStat{
									Started: 2, Ended: 2, Count: 1,
									Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
									Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
									Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
									SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
									Success:      0,
									SuccessCodes: map[string]uint64{},
									ErrorCodes:   map[string]uint64{"500": 1},
								},
							},
							SumStat: aggstat.AggStat{
								Started: 1, Ended: 2, Count: 2,
								Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
								Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
								Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
								SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
								Success:      1,
								SuccessCodes: map[string]uint64{"200": 1},
								ErrorCodes:   map[string]uint64{"500": 2},
							},
						},
					},
				},
				o: aggstat.LabelURLAggStat{Started: 1, Ended: 3, Name: "Test (Cmp)",
					Stat: map[string]*aggstat.URLAggStat{
						"test1": {
							Stat: map[string]*aggstat.AggStat{
								"/test1": &aggstat.AggStat{
									Started: 1, Ended: 1, Count: 1,
									Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
									Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
									Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
									SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
									Success:      1,
									SuccessCodes: map[string]uint64{"200": 1},
									ErrorCodes:   map[string]uint64{},
								},
								"/test2": &aggstat.AggStat{
									Started: 3, Ended: 3, Count: 1,
									Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
									Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
									Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
									SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
									Success:      0,
									SuccessCodes: map[string]uint64{},
									ErrorCodes:   map[string]uint64{"500": 1},
								},
							},
							SumStat: aggstat.AggStat{
								Started: 1, Ended: 3, Count: 2,
								Elapsed:      aggstat.AggStatNode{Min: 10, Max: 10, Mean: 10, P90: 10, P95: 10, P99: 10},
								Connect:      aggstat.AggStatNode{Min: 1, Max: 1, Mean: 1, P90: 1, P95: 1, P99: 1},
								Bytes:        aggstat.AggStatNode{Min: 200, Max: 200, Mean: 200, P90: 200, P95: 200, P99: 200},
								SentBytes:    aggstat.AggStatNode{Min: 20, Max: 20, Mean: 20, P90: 20, P95: 20, P99: 20},
								Success:      1,
								SuccessCodes: map[string]uint64{"200": 1},
								ErrorCodes:   map[string]uint64{"500": 1},
							},
						},
					},
				},
			},
			LabelURLAggDiffStat{
				Name: "Test", Started: 1, Ended: 2,
				CmpName: "Test (Cmp)", CmpStarted: 1, CmpEnded: 3,
				Stat: map[string]*URLAggDiffStat{
					"test1": &URLAggDiffStat{
						Stat: map[string]*AggDiffStat{
							"/test1": &AggDiffStat{
								Elapsed:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Connect:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Bytes:        aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								SentBytes:    aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Errors:       50.0,
								SuccessCodes: map[string]float64{"200": -50.0},
								ErrorCodes:   map[string]float64{"500": 50.0},
							},
							"/test2": &AggDiffStat{
								Elapsed:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Connect:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Bytes:        aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								SentBytes:    aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
								Errors:       0.0,
								SuccessCodes: map[string]float64{},
								ErrorCodes:   map[string]float64{"500": 0.0},
							},
						},
						SumStat: AggDiffStat{
							Elapsed:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
							Connect:      aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
							Bytes:        aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
							SentBytes:    aggstat.AggStatNode{Min: 0.0, Max: 0.0, Mean: 0.0, P90: 0.0, P95: 0.0, P99: 0.0},
							Errors:       16.666666666666664,
							SuccessCodes: map[string]float64{"200": -16.666666666666668},
							ErrorCodes:   map[string]float64{"500": 16.666666666666664},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		var got LabelURLAggDiffStat
		got.DiffPcnt(&tt.args.s, &tt.args.o)
		if !reflect.DeepEqual(got, tt.result) {
			t.Errorf("LabelURLAggDiffStat.DiffPcnt() =\n%# v\n, want\n%# v\n, diff\n%v", pretty.Formatter(got), pretty.Formatter(tt.result), pretty.Diff(tt.result, got))
		}
	}
}
