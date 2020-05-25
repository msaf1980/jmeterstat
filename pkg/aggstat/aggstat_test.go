package aggstat

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
)

func TestLabelURLAggStat(t *testing.T) {
	tests := []struct {
		name   string
		input  jmeterstat.JMeterLabelURLStat
		result LabelURLAggStat
	}{
		{
			"/test1 + /test2",
			jmeterstat.JMeterLabelURLStat{
				"test1": {
					"/test1": new(jmeterstat.JMeterStat).Init().Add(1589308021000, 2.0, 1.0, 10.0, 120.0, true, "200").
						Add(1589308022000, 2.0, 1.0, 10.0, 120.0, false, "500"),
					"/test11": new(jmeterstat.JMeterStat).Init().Add(1589308023000, 2.0, 1.0, 10.0, 120.0, true, "200"),
				},
				"test2": {
					"/test2": new(jmeterstat.JMeterStat).Init().Add(1589308023000, 2.0, 1.0, 10.0, 120.0, true, "200"),
				},
			},
			LabelURLAggStat{
				Started: 1589308021000,
				Ended:   1589308023000,
				Name:    "test",
				Stat: map[string]*URLAggStat{
					"test1": &URLAggStat{
						Stat: map[string]*AggStat{
							"/test1": &AggStat{
								Started: 1589308021000, Ended: 1589308022000, Count: 2,
								Elapsed: AggStatNode{
									Min: 2.0, Max: 2.0, Mean: 2.0, P90: 2.0, P95: 2.0, P99: 2.0,
								},
								Connect: AggStatNode{
									Min: 1.0, Max: 1.0, Mean: 1.0, P90: 1.0, P95: 1.0, P99: 1.0,
								},
								Bytes: AggStatNode{
									Min: 10.0, Max: 10.0, Mean: 10.0, P90: 10.0, P95: 10.0, P99: 10.0,
								},
								SentBytes: AggStatNode{
									Min: 120.0, Max: 120.0, Mean: 120.0, P90: 120.0, P95: 120.0, P99: 120.0,
								},
								Success:      1,
								SuccessCodes: map[string]uint64{"200": 1},
								ErrorCodes:   map[string]uint64{"500": 1},
							},
							"/test11": &AggStat{
								Started: 1589308023000, Ended: 1589308023000, Count: 1,
								Elapsed: AggStatNode{
									Min: 2.0, Max: 2.0, Mean: 2.0, P90: 2.0, P95: 2.0, P99: 2.0,
								},
								Connect: AggStatNode{
									Min: 1.0, Max: 1.0, Mean: 1.0, P90: 1.0, P95: 1.0, P99: 1.0,
								},
								Bytes: AggStatNode{
									Min: 10.0, Max: 10.0, Mean: 10.0, P90: 10.0, P95: 10.0, P99: 10.0,
								},
								SentBytes: AggStatNode{
									Min: 120.0, Max: 120.0, Mean: 120.0, P90: 120.0, P95: 120.0, P99: 120.0,
								},
								Success:      1,
								SuccessCodes: map[string]uint64{"200": 1},
								ErrorCodes:   map[string]uint64{},
							},
						},
						SumStat: AggStat{
							Started: 1589308021000, Ended: 1589308023000, Count: 3,
							Elapsed: AggStatNode{
								Min: 2.0, Max: 2.0, Mean: 2.0, P90: 2.0, P95: 2.0, P99: 2.0,
							},
							Connect: AggStatNode{
								Min: 1.0, Max: 1.0, Mean: 1.0, P90: 1.0, P95: 1.0, P99: 1.0,
							},
							Bytes: AggStatNode{
								Min: 10.0, Max: 10.0, Mean: 10.0, P90: 10.0, P95: 10.0, P99: 10.0,
							},
							SentBytes: AggStatNode{
								Min: 120.0, Max: 120.0, Mean: 120.0, P90: 120.0, P95: 120.0, P99: 120.0,
							},
							Success:      2,
							SuccessCodes: map[string]uint64{"200": 2},
							ErrorCodes:   map[string]uint64{"500": 1},
						},
					},
					/*
						"test2": {
							"/test2": new(jmeterstat.JMeterStat).Init().Add(3, 2.0, 1.0, 10.0, 120.0, true, "200"),
						},
					*/
					"test2": &URLAggStat{
						Stat: map[string]*AggStat{
							"/test2": &AggStat{
								Started: 1589308023000, Ended: 1589308023000, Count: 1,
								Elapsed: AggStatNode{
									Min: 2.0, Max: 2.0, Mean: 2.0, P90: 2.0, P95: 2.0, P99: 2.0,
								},
								Connect: AggStatNode{
									Min: 1.0, Max: 1.0, Mean: 1.0, P90: 1.0, P95: 1.0, P99: 1.0,
								},
								Bytes: AggStatNode{
									Min: 10.0, Max: 10.0, Mean: 10.0, P90: 10.0, P95: 10.0, P99: 10.0,
								},
								SentBytes: AggStatNode{
									Min: 120.0, Max: 120.0, Mean: 120.0, P90: 120.0, P95: 120.0, P99: 120.0,
								},
								Success:      1,
								SuccessCodes: map[string]uint64{"200": 1},
								ErrorCodes:   map[string]uint64{},
							},
						},
						SumStat: AggStat{
							Started: 1589308023000, Ended: 1589308023000, Count: 1,
							Elapsed: AggStatNode{
								Min: 2.0, Max: 2.0, Mean: 2.0, P90: 2.0, P95: 2.0, P99: 2.0,
							},
							Connect: AggStatNode{
								Min: 1.0, Max: 1.0, Mean: 1.0, P90: 1.0, P95: 1.0, P99: 1.0,
							},
							Bytes: AggStatNode{
								Min: 10.0, Max: 10.0, Mean: 10.0, P90: 10.0, P95: 10.0, P99: 10.0,
							},
							SentBytes: AggStatNode{
								Min: 120.0, Max: 120.0, Mean: 120.0, P90: 120.0, P95: 120.0, P99: 120.0,
							},
							Success:      1,
							SuccessCodes: map[string]uint64{"200": 1},
							ErrorCodes:   map[string]uint64{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		var got LabelURLAggStat
		got.Init(tt.input, "test")
		if !reflect.DeepEqual(got, tt.result) {
			t.Errorf("LabelURLAggStat.Init() =\n%# v\n, want\n%# v\n, diff\n%v", pretty.Formatter(got), pretty.Formatter(tt.result), pretty.Diff(tt.result, got))
		}

		// Test JSON Marshal/UnMarshal
		obytes, err := got.MarshalJSON()
		if err != nil {
			t.Errorf("LabelURLAggStat.MarshalJSON() err: %s", err.Error())
		}
		var restore LabelURLAggStat
		err = restore.UnmarshalJSON(obytes)
		if err != nil {
			t.Errorf("LabelURLAggStat.UnMarshalJSON() err: %s", err.Error())
		}
		if !reflect.DeepEqual(restore, tt.result) {
			t.Errorf("LabelURLAggStat..UnMarshalJSON() =\n%# v\n, want\n%# v\n, diff\n%v", pretty.Formatter(restore), pretty.Formatter(tt.result), pretty.Diff(tt.result, restore))
		}
	}
}
