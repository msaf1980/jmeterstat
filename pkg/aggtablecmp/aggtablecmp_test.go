package aggtablecmp

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func EqualRequests(a, b []RequestDiffStat) bool {
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

func TestRequestsStat_Sort(t *testing.T) {
	type args struct {
		sortCol  SortColumn
		sortDesc bool
	}
	input := RequestsDiffStat{
		Label: "test",
		Stat: []RequestDiffStat{
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
				Errors:  -6,

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
	}
	tests := []struct {
		name   string
		args   args
		result []RequestDiffStat
	}{
		{
			"Requests (asc)",
			args{SortRequest, false},
			[]RequestDiffStat{
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
					Errors:  -6,

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
			[]RequestDiffStat{
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
					Errors:  -6,

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
	}

	for _, tt := range tests {
		input.SortRequests(tt.args.sortCol, tt.args.sortDesc)
		if !EqualRequests(input.Stat, tt.result) {
			t.Errorf("%s, got\n%# v\n, want\n%# v\n", tt.name, pretty.Formatter(input.Stat), pretty.Formatter(tt.result))
		}
	}
}
