package aggtable

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/msaf1980/jmeterstat/pkg/aggstat"
)

func equalRequests(a, b []RequestStat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalErrors(a, b []ErrorStat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
		if len(a[i].ErrorCodes) != len(b[i].ErrorCodes) {
			return false
		}

		for j := range a[i].ErrorCodes {
			if a[i].ErrorCodes[j] != b[i].ErrorCodes[j] {
				return false
			}
		}
	}
	return true
}

func TestRequestsStat_SortRequests(t *testing.T) {
	type args struct {
		sortCol  SortColumn
		sortDesc bool
	}
	input := RequestsStat{
		Label: "test",
		Stat: []RequestStat{
			{
				Request: "/test1",
				Samples: 10,
				Errors:  5,

				ResponceTimeMean: 2,
				ResponceTimeMin:  1,
				ResponceTimeMax:  10,
				ResponceTimeP90:  3,
				ResponceTimeP95:  4,
				ResponceTimeP99:  5,

				SentMean: 3,
				SentMin:  2,
				SentMax:  11,
				SentP90:  4,
				SentP95:  5,
				SentP99:  6,

				ReceivedMean: 4,
				ReceivedMin:  3,
				ReceivedMax:  12,
				ReceivedP90:  5,
				ReceivedP95:  6,
				ReceivedP99:  7,
			},
			{
				Request: "/test3",
				Samples: 4,
				Errors:  2,

				ResponceTimeMean: 20,
				ResponceTimeMin:  10,
				ResponceTimeMax:  100,
				ResponceTimeP90:  30,
				ResponceTimeP95:  40,
				ResponceTimeP99:  50,

				SentMean: 30,
				SentMin:  20,
				SentMax:  110,
				SentP90:  40,
				SentP95:  50,
				SentP99:  60,

				ReceivedMean: 40,
				ReceivedMin:  30,
				ReceivedMax:  120,
				ReceivedP90:  50,
				ReceivedP95:  60,
				ReceivedP99:  70,
			},
			{
				Request: "/test2",
				Samples: 8,
				Errors:  6,

				ResponceTimeMean: 4,
				ResponceTimeMin:  2,
				ResponceTimeMax:  20,
				ResponceTimeP90:  6,
				ResponceTimeP95:  8,
				ResponceTimeP99:  10,

				SentMean: 6,
				SentMin:  4,
				SentMax:  22,
				SentP90:  8,
				SentP95:  10,
				SentP99:  12,

				ReceivedMean: 8,
				ReceivedMin:  6,
				ReceivedMax:  24,
				ReceivedP90:  10,
				ReceivedP95:  12,
				ReceivedP99:  14,
			},
		},
		ErrStat: []ErrorStat{},
	}
	tests := []struct {
		name   string
		args   args
		result []RequestStat
	}{
		{
			"Requests (asc)",
			args{SortRequest, false},
			[]RequestStat{
				{
					Request: "/test1",
					Samples: 10,
					Errors:  5,

					ResponceTimeMean: 2,
					ResponceTimeMin:  1,
					ResponceTimeMax:  10,
					ResponceTimeP90:  3,
					ResponceTimeP95:  4,
					ResponceTimeP99:  5,

					SentMean: 3,
					SentMin:  2,
					SentMax:  11,
					SentP90:  4,
					SentP95:  5,
					SentP99:  6,

					ReceivedMean: 4,
					ReceivedMin:  3,
					ReceivedMax:  12,
					ReceivedP90:  5,
					ReceivedP95:  6,
					ReceivedP99:  7,
				},
				{
					Request: "/test2",
					Samples: 8,
					Errors:  6,

					ResponceTimeMean: 4,
					ResponceTimeMin:  2,
					ResponceTimeMax:  20,
					ResponceTimeP90:  6,
					ResponceTimeP95:  8,
					ResponceTimeP99:  10,

					SentMean: 6,
					SentMin:  4,
					SentMax:  22,
					SentP90:  8,
					SentP95:  10,
					SentP99:  12,

					ReceivedMean: 8,
					ReceivedMin:  6,
					ReceivedMax:  24,
					ReceivedP90:  10,
					ReceivedP95:  12,
					ReceivedP99:  14,
				},
				{
					Request: "/test3",
					Samples: 4,
					Errors:  2,

					ResponceTimeMean: 20,
					ResponceTimeMin:  10,
					ResponceTimeMax:  100,
					ResponceTimeP90:  30,
					ResponceTimeP95:  40,
					ResponceTimeP99:  50,

					SentMean: 30,
					SentMin:  20,
					SentMax:  110,
					SentP90:  40,
					SentP95:  50,
					SentP99:  60,

					ReceivedMean: 40,
					ReceivedMin:  30,
					ReceivedMax:  120,
					ReceivedP90:  50,
					ReceivedP95:  60,
					ReceivedP99:  70,
				},
			},
		},
		{
			"Requests (desc)",
			args{SortRequest, true},
			[]RequestStat{
				{
					Request: "/test3",
					Samples: 4,
					Errors:  2,

					ResponceTimeMean: 20,
					ResponceTimeMin:  10,
					ResponceTimeMax:  100,
					ResponceTimeP90:  30,
					ResponceTimeP95:  40,
					ResponceTimeP99:  50,

					SentMean: 30,
					SentMin:  20,
					SentMax:  110,
					SentP90:  40,
					SentP95:  50,
					SentP99:  60,

					ReceivedMean: 40,
					ReceivedMin:  30,
					ReceivedMax:  120,
					ReceivedP90:  50,
					ReceivedP95:  60,
					ReceivedP99:  70,
				},
				{
					Request: "/test2",
					Samples: 8,
					Errors:  6,

					ResponceTimeMean: 4,
					ResponceTimeMin:  2,
					ResponceTimeMax:  20,
					ResponceTimeP90:  6,
					ResponceTimeP95:  8,
					ResponceTimeP99:  10,

					SentMean: 6,
					SentMin:  4,
					SentMax:  22,
					SentP90:  8,
					SentP95:  10,
					SentP99:  12,

					ReceivedMean: 8,
					ReceivedMin:  6,
					ReceivedMax:  24,
					ReceivedP90:  10,
					ReceivedP95:  12,
					ReceivedP99:  14,
				},
				{
					Request: "/test1",
					Samples: 10,
					Errors:  5,

					ResponceTimeMean: 2,
					ResponceTimeMin:  1,
					ResponceTimeMax:  10,
					ResponceTimeP90:  3,
					ResponceTimeP95:  4,
					ResponceTimeP99:  5,

					SentMean: 3,
					SentMin:  2,
					SentMax:  11,
					SentP90:  4,
					SentP95:  5,
					SentP99:  6,

					ReceivedMean: 4,
					ReceivedMin:  3,
					ReceivedMax:  12,
					ReceivedP90:  5,
					ReceivedP95:  6,
					ReceivedP99:  7,
				},
			},
		},
		{
			"Errors (asc)",
			args{SortErrors, false},
			[]RequestStat{
				{
					Request: "/test3",
					Samples: 4,
					Errors:  2,

					ResponceTimeMean: 20,
					ResponceTimeMin:  10,
					ResponceTimeMax:  100,
					ResponceTimeP90:  30,
					ResponceTimeP95:  40,
					ResponceTimeP99:  50,

					SentMean: 30,
					SentMin:  20,
					SentMax:  110,
					SentP90:  40,
					SentP95:  50,
					SentP99:  60,

					ReceivedMean: 40,
					ReceivedMin:  30,
					ReceivedMax:  120,
					ReceivedP90:  50,
					ReceivedP95:  60,
					ReceivedP99:  70,
				},
				{
					Request: "/test1",
					Samples: 10,
					Errors:  5,

					ResponceTimeMean: 2,
					ResponceTimeMin:  1,
					ResponceTimeMax:  10,
					ResponceTimeP90:  3,
					ResponceTimeP95:  4,
					ResponceTimeP99:  5,

					SentMean: 3,
					SentMin:  2,
					SentMax:  11,
					SentP90:  4,
					SentP95:  5,
					SentP99:  6,

					ReceivedMean: 4,
					ReceivedMin:  3,
					ReceivedMax:  12,
					ReceivedP90:  5,
					ReceivedP95:  6,
					ReceivedP99:  7,
				},
				{
					Request: "/test2",
					Samples: 8,
					Errors:  6,

					ResponceTimeMean: 4,
					ResponceTimeMin:  2,
					ResponceTimeMax:  20,
					ResponceTimeP90:  6,
					ResponceTimeP95:  8,
					ResponceTimeP99:  10,

					SentMean: 6,
					SentMin:  4,
					SentMax:  22,
					SentP90:  8,
					SentP95:  10,
					SentP99:  12,

					ReceivedMean: 8,
					ReceivedMin:  6,
					ReceivedMax:  24,
					ReceivedP90:  10,
					ReceivedP95:  12,
					ReceivedP99:  14,
				},
			},
		},
		{
			"Errors (desc)",
			args{SortErrors, true},
			[]RequestStat{
				{
					Request: "/test2",
					Samples: 8,
					Errors:  6,

					ResponceTimeMean: 4,
					ResponceTimeMin:  2,
					ResponceTimeMax:  20,
					ResponceTimeP90:  6,
					ResponceTimeP95:  8,
					ResponceTimeP99:  10,

					SentMean: 6,
					SentMin:  4,
					SentMax:  22,
					SentP90:  8,
					SentP95:  10,
					SentP99:  12,

					ReceivedMean: 8,
					ReceivedMin:  6,
					ReceivedMax:  24,
					ReceivedP90:  10,
					ReceivedP95:  12,
					ReceivedP99:  14,
				},
				{
					Request: "/test1",
					Samples: 10,
					Errors:  5,

					ResponceTimeMean: 2,
					ResponceTimeMin:  1,
					ResponceTimeMax:  10,
					ResponceTimeP90:  3,
					ResponceTimeP95:  4,
					ResponceTimeP99:  5,

					SentMean: 3,
					SentMin:  2,
					SentMax:  11,
					SentP90:  4,
					SentP95:  5,
					SentP99:  6,

					ReceivedMean: 4,
					ReceivedMin:  3,
					ReceivedMax:  12,
					ReceivedP90:  5,
					ReceivedP95:  6,
					ReceivedP99:  7,
				},
				{
					Request: "/test3",
					Samples: 4,
					Errors:  2,

					ResponceTimeMean: 20,
					ResponceTimeMin:  10,
					ResponceTimeMax:  100,
					ResponceTimeP90:  30,
					ResponceTimeP95:  40,
					ResponceTimeP99:  50,

					SentMean: 30,
					SentMin:  20,
					SentMax:  110,
					SentP90:  40,
					SentP95:  50,
					SentP99:  60,

					ReceivedMean: 40,
					ReceivedMin:  30,
					ReceivedMax:  120,
					ReceivedP90:  50,
					ReceivedP95:  60,
					ReceivedP99:  70,
				},
			},
		},
	}

	for _, tt := range tests {
		input.SortRequests(tt.args.sortCol, tt.args.sortDesc)
		if !equalRequests(input.Stat, tt.result) {
			t.Errorf("%s, got\n%# v\n, want\n%# v\n", tt.name, pretty.Formatter(input.Stat), pretty.Formatter(tt.result))
		}
	}
}

func TestRequestStat_Init(t *testing.T) {
	type args struct {
		request string
		stat    aggstat.AggStat
	}
	tests := []struct {
		name   string
		args   args
		result RequestStat
	}{
		{
			name: "errors extend",
			args: args{
				request: "/test",
				stat: aggstat.AggStat{
					Started:      10,
					Ended:        20,
					Count:        5,
					Elapsed:      aggstat.AggStatNode{Min: 1, Max: 3, Mean: 2, P90: 2, P95: 2, P99: 2},
					Connect:      aggstat.AggStatNode{Min: 10, Max: 30, Mean: 20, P90: 20, P95: 20, P99: 20},
					Bytes:        aggstat.AggStatNode{Min: 10, Max: 30, Mean: 20, P90: 20, P95: 20, P99: 20},
					SentBytes:    aggstat.AggStatNode{Min: 1, Max: 3, Mean: 2, P90: 2, P95: 2, P99: 2},
					Success:      2,
					SuccessCodes: map[string]uint64{"200": 2},
					ErrorCodes:   map[string]uint64{"500": 1, "504": 2},
				},
			},
			result: RequestStat{
				Request:          "/test",
				Samples:          5,
				Errors:           60.0,
				ResponceTimeMean: 2.0,
				ResponceTimeMin:  1.0,
				ResponceTimeMax:  3.0,
				ResponceTimeP90:  2.0,
				ResponceTimeP95:  2.0,
				ResponceTimeP99:  2.0,
				SentMean:         2.0,
				SentMin:          1.0,
				SentMax:          3.0,
				SentP90:          2.0,
				SentP95:          2.0,
				SentP99:          2.0,
				ReceivedMean:     20.0,
				ReceivedMin:      10.0,
				ReceivedMax:      30.0,
				ReceivedP90:      20.0,
				ReceivedP95:      20.0,
				ReceivedP99:      20.0,
			},
		},
		{
			name: "errors truncated",
			args: args{
				request: "/test",
				stat: aggstat.AggStat{
					Started:      10,
					Ended:        20,
					Count:        10,
					Elapsed:      aggstat.AggStatNode{Min: 1, Max: 3, Mean: 2, P90: 2, P95: 2, P99: 2},
					Connect:      aggstat.AggStatNode{Min: 10, Max: 30, Mean: 20, P90: 20, P95: 20, P99: 20},
					Bytes:        aggstat.AggStatNode{Min: 10, Max: 30, Mean: 20, P90: 20, P95: 20, P99: 20},
					SentBytes:    aggstat.AggStatNode{Min: 1, Max: 3, Mean: 2, P90: 2, P95: 2, P99: 2},
					Success:      2,
					SuccessCodes: map[string]uint64{"200": 2},
					ErrorCodes:   map[string]uint64{"500": 1, "403": 2, "504": 2, "502": 5},
				},
			},
			result: RequestStat{
				Request:          "/test",
				Samples:          10,
				Errors:           80.0,
				ResponceTimeMean: 2.0,
				ResponceTimeMin:  1.0,
				ResponceTimeMax:  3.0,
				ResponceTimeP90:  2.0,
				ResponceTimeP95:  2.0,
				ResponceTimeP99:  2.0,
				SentMean:         2.0,
				SentMin:          1.0,
				SentMax:          3.0,
				SentP90:          2.0,
				SentP95:          2.0,
				SentP99:          2.0,
				ReceivedMean:     20.0,
				ReceivedMin:      10.0,
				ReceivedMax:      30.0,
				ReceivedP90:      20.0,
				ReceivedP95:      20.0,
				ReceivedP99:      20.0,
			},
		},
	}
	for _, tt := range tests {
		r := RequestStat{}
		r.Init(tt.args.request, &tt.args.stat)
		if !reflect.DeepEqual(r, tt.result) {
			t.Errorf("%s, got\n%# v\n, want\n%# v\n, diff\n%# v\n", tt.name, pretty.Formatter(r), pretty.Formatter(tt.result), pretty.Diff(tt.result, r))
		}
	}
}

func TestRequestsStat_SortErrors(t *testing.T) {
	input := []ErrorStat{
		{
			Request: "/test2",
			Samples: 16,
			ErrorCodes: []ErrorCount{
				{"500", 1},
				{"401", 3},
				{"503", 5},
				{"504", 2},
				{"403", 1},
			},
		},
		{
			Request: "/test1",
			Samples: 12,
			ErrorCodes: []ErrorCount{
				{"502", 5},
				{"403", 2},
				{"504", 2},
				{"", 0},
				{"", 0},
			},
		},
		{
			Request: "/test3",
			Samples: 20,
			ErrorCodes: []ErrorCount{
				{"500", 2},
				{"401", 1},
				{"503", 6},
				{"504", 1},
				{"403", 1},
			},
		},
	}

	type args struct {
		sortCol  int
		sortDesc bool
	}
	tests := []struct {
		name   string
		args   args
		result []ErrorStat
	}{
		{
			"SortErrors(0, false)",
			args{0, false},
			[]ErrorStat{
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
			},
		},
		{
			"SortErrors(0, true)",
			args{0, true},
			[]ErrorStat{
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
			},
		},

		{
			"SortErrors(1, false)",
			args{1, false},
			[]ErrorStat{
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
			},
		},

		{
			"SortErrors(1, true)",
			args{1, true},
			[]ErrorStat{
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
			},
		},
		{
			"SortErrors(2, false)",
			args{2, false},
			[]ErrorStat{
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
			},
		},

		{
			"SortErrors(2, true)",
			args{2, true},
			[]ErrorStat{
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
			},
		},

		{
			"SortErrors(3, false)",
			args{3, false},
			[]ErrorStat{
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
			},
		},

		{
			"SortErrors(4, true)",
			args{4, true},
			[]ErrorStat{
				{
					Request: "/test2",
					Samples: 16,
					ErrorCodes: []ErrorCount{
						{"500", 1},
						{"401", 3},
						{"503", 5},
						{"504", 2},
						{"403", 1},
					},
				},
				{
					Request: "/test1",
					Samples: 12,
					ErrorCodes: []ErrorCount{
						{"502", 5},
						{"403", 2},
						{"504", 2},
						{"", 0},
						{"", 0},
					},
				},
				{
					Request: "/test3",
					Samples: 20,
					ErrorCodes: []ErrorCount{
						{"500", 2},
						{"401", 1},
						{"503", 6},
						{"504", 1},
						{"403", 1},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RequestsStat{ErrStat: input}
			r.SortErrors(tt.args.sortCol, tt.args.sortDesc)
			if !equalErrors(r.ErrStat, tt.result) {
				t.Errorf("%s, got\n%# v\n, want\n%# v\n, diff\n%# v\n", tt.name, pretty.Formatter(r.ErrStat), pretty.Formatter(tt.result), pretty.Diff(r.ErrStat, tt.result))
			}
		})
	}
}
