package aggstat

import (
	"reflect"
	"testing"

	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
)

func TestLabelURLAggStat(t *testing.T) {
	tests := []struct {
		name    string
		input   jmeterstat.JMeterLabelURLStat
		result  LabelURLAggStat
		jsonStr string
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
								Success:       1,
								ResponceCodes: map[string]uint64{"200": 1, "500": 1},
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
								Success:       1,
								ResponceCodes: map[string]uint64{"200": 1},
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
							Success:       2,
							ResponceCodes: map[string]uint64{"200": 2, "500": 1},
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
								Success:       1,
								ResponceCodes: map[string]uint64{"200": 1},
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
							Success:       1,
							ResponceCodes: map[string]uint64{"200": 1},
						},
					},
				},
			},
			"{\"Started\":1589308021000,\"Ended\":1589308023000,\"Stat\":{\"test1\":{\"Stat\":{\"/test1\":{\"Started\":1589308021000,\"Ended\":1589308022000,\"Count\":2,\"Elapsed\":{\"Min\":2,\"Max\":2,\"Mean\":2,\"P90\":2,\"P95\":2,\"P99\":2},\"Connect\":{\"Min\":1,\"Max\":1,\"Mean\":1,\"P90\":1,\"P95\":1,\"P99\":1},\"Bytes\":{\"Min\":10,\"Max\":10,\"Mean\":10,\"P90\":10,\"P95\":10,\"P99\":10},\"SentBytes\":{\"Min\":120,\"Max\":120,\"Mean\":120,\"P90\":120,\"P95\":120,\"P99\":120},\"Success\":1,\"ResponceCodes\":{\"200\":1,\"500\":1}},\"/test11\":{\"Started\":1589308023000,\"Ended\":1589308023000,\"Count\":1,\"Elapsed\":{\"Min\":2,\"Max\":2,\"Mean\":2,\"P90\":2,\"P95\":2,\"P99\":2},\"Connect\":{\"Min\":1,\"Max\":1,\"Mean\":1,\"P90\":1,\"P95\":1,\"P99\":1},\"Bytes\":{\"Min\":10,\"Max\":10,\"Mean\":10,\"P90\":10,\"P95\":10,\"P99\":10},\"SentBytes\":{\"Min\":120,\"Max\":120,\"Mean\":120,\"P90\":120,\"P95\":120,\"P99\":120},\"Success\":1,\"ResponceCodes\":{\"200\":1}}},\"SumStat\":{\"Started\":1589308021000,\"Ended\":1589308023000,\"Count\":3,\"Elapsed\":{\"Min\":2,\"Max\":2,\"Mean\":2,\"P90\":2,\"P95\":2,\"P99\":2},\"Connect\":{\"Min\":1,\"Max\":1,\"Mean\":1,\"P90\":1,\"P95\":1,\"P99\":1},\"Bytes\":{\"Min\":10,\"Max\":10,\"Mean\":10,\"P90\":10,\"P95\":10,\"P99\":10},\"SentBytes\":{\"Min\":120,\"Max\":120,\"Mean\":120,\"P90\":120,\"P95\":120,\"P99\":120},\"Success\":2,\"ResponceCodes\":{\"200\":2,\"500\":1}}},\"test2\":{\"Stat\":{\"/test2\":{\"Started\":1589308023000,\"Ended\":1589308023000,\"Count\":1,\"Elapsed\":{\"Min\":2,\"Max\":2,\"Mean\":2,\"P90\":2,\"P95\":2,\"P99\":2},\"Connect\":{\"Min\":1,\"Max\":1,\"Mean\":1,\"P90\":1,\"P95\":1,\"P99\":1},\"Bytes\":{\"Min\":10,\"Max\":10,\"Mean\":10,\"P90\":10,\"P95\":10,\"P99\":10},\"SentBytes\":{\"Min\":120,\"Max\":120,\"Mean\":120,\"P90\":120,\"P95\":120,\"P99\":120},\"Success\":1,\"ResponceCodes\":{\"200\":1}}},\"SumStat\":{\"Started\":1589308023000,\"Ended\":1589308023000,\"Count\":1,\"Elapsed\":{\"Min\":2,\"Max\":2,\"Mean\":2,\"P90\":2,\"P95\":2,\"P99\":2},\"Connect\":{\"Min\":1,\"Max\":1,\"Mean\":1,\"P90\":1,\"P95\":1,\"P99\":1},\"Bytes\":{\"Min\":10,\"Max\":10,\"Mean\":10,\"P90\":10,\"P95\":10,\"P99\":10},\"SentBytes\":{\"Min\":120,\"Max\":120,\"Mean\":120,\"P90\":120,\"P95\":120,\"P99\":120},\"Success\":1,\"ResponceCodes\":{\"200\":1}}}}}",
		},
	}
	for _, tt := range tests {
		var got LabelURLAggStat
		got.Init(tt.input, "test")
		if !reflect.DeepEqual(got, tt.result) {
			t.Errorf("LabelURLAggStat.Init() = %v, want %v", got, tt.result)
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
			t.Errorf("LabelURLAggStat.UnMarshalJSON() = %v, want %v", got, tt.result)
		}
	}
}
