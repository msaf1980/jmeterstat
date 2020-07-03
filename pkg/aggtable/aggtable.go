package aggtable

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
)

// SortColumn sort by column
type SortColumn int

const (
	// Unsorted unsorted
	Unsorted SortColumn = iota
	// SortRequest sort by request name
	SortRequest
	// SortSamples sort by samples count
	SortSamples
	// SortErrors sort by errors count
	SortErrors
	// SortResponceTimeMean sort by responce time mean
	SortResponceTimeMean
	// SortResponceTimeMin sort by responce time min
	SortResponceTimeMin
	// SortResponceTimeMax sort by responce time max
	SortResponceTimeMax
	// SortResponceTimeP90 sort by responce time P90
	SortResponceTimeP90
	// SortResponceTimeP95 sort by responce time P95
	SortResponceTimeP95
	// SortResponceTimeP99 sort by responce time P99
	SortResponceTimeP99
	// SortSentMean sort by sent bytes mean
	SortSentMean
	// SortSentMin sort by sent bytes min
	SortSentMin
	// SortSentMax sort by sent bytes max
	SortSentMax
	// SortSentP90 sort by sent bytes P90
	SortSentP90
	// SortSentP95 sort by sent bytes P95
	SortSentP95
	// SortSentP99 sort by sent bytes P99
	SortSentP99
	// SortReceivedMean sort by received bytes mean
	SortReceivedMean
	// SortReceivedMin sort by received bytes min
	SortReceivedMin
	// SortReceivedMax sort by received bytes max
	SortReceivedMax
	// SortReceivedP90 sort by received bytes P90
	SortReceivedP90
	// SortReceivedP95 sort by received bytes P95
	SortReceivedP95
	// SortReceivedP99 sort by received bytes P99
	SortReceivedP99
)

// GetSortColumn get column sort
func GetSortColumn(sortColumn string) (SortColumn, error) {
	s := Unsorted

	if n, err := strconv.Atoi(sortColumn); err != nil {
		return s, fmt.Errorf("Uncorrect order %s", sortColumn)
	} else {
		switch n {
		case 0:
			s = SortRequest
		case 1:
			s = SortSamples
		case 2:
			s = SortErrors
		case 3:
			s = SortResponceTimeMean
		case 4:
			s = SortResponceTimeMin
		case 5:
			s = SortResponceTimeMax
		case 6:
			s = SortResponceTimeP90
		case 7:
			s = SortResponceTimeP95
		case 8:
			s = SortResponceTimeP99
		case 9:
			s = SortSentMean
		case 10:
			s = SortSentMin
		case 11:
			s = SortSentMax
		case 12:
			s = SortSentP90
		case 13:
			s = SortSentP95
		case 14:
			s = SortSentP99
		case 15:
			s = SortReceivedMean
		case 16:
			s = SortReceivedMin
		case 17:
			s = SortReceivedMax
		case 18:
			s = SortReceivedP90
		case 19:
			s = SortReceivedP95
		case 20:
			s = SortReceivedP99
		default:
			return s, fmt.Errorf("Unknown order %s", sortColumn)
		}
		return s, nil
	}
}

// ErrorCount error count
type ErrorCount struct {
	Error  string
	Errors uint64
}

// RequestStat for request statictic
type RequestStat struct {
	Request string

	Samples uint64
	Errors  float64

	ResponceTimeMean float64
	ResponceTimeMin  float64
	ResponceTimeMax  float64
	ResponceTimeP90  float64
	ResponceTimeP95  float64
	ResponceTimeP99  float64

	SentMean float64
	SentMin  float64
	SentMax  float64
	SentP90  float64
	SentP95  float64
	SentP99  float64

	ReceivedMean float64
	ReceivedMin  float64
	ReceivedMax  float64
	ReceivedP90  float64
	ReceivedP95  float64
	ReceivedP99  float64

	ErrorCodes []ErrorCount
}

func maxErrors(errCount map[string]uint64, maxCount int) []ErrorCount {
	err := make([]ErrorCount, len(errCount))
	i := 0
	for code := range errCount {
		err[i] = ErrorCount{Error: code, Errors: errCount[code]}
		i++
	}
	sort.Slice(err, func(i, j int) bool {
		if err[i].Errors == err[j].Errors {
			return strings.Compare(err[i].Error, err[j].Error) < 0
		}
		return err[i].Errors > err[j].Errors
	})

	errDst := make([]ErrorCount, maxCount)
	length := maxCount
	if length > len(errCount) {
		length = len(errCount)
	}
	copy(errDst, err[0:length])

	return errDst
}

// Init from aggstat.AggStat
func (r *RequestStat) Init(request string, stat *aggstat.AggStat, errors int) {
	r.Request = request

	r.Samples = stat.Count
	r.Errors = 100 * (float64(stat.Count) - float64(stat.Success)) / float64(stat.Count)

	r.ResponceTimeMean = stat.Elapsed.Mean
	r.ResponceTimeMin = stat.Elapsed.Min
	r.ResponceTimeMax = stat.Elapsed.Max
	r.ResponceTimeP90 = stat.Elapsed.P90
	r.ResponceTimeP95 = stat.Elapsed.P95
	r.ResponceTimeP99 = stat.Elapsed.P99

	r.SentMean = stat.SentBytes.Mean
	r.SentMin = stat.SentBytes.Min
	r.SentMax = stat.SentBytes.Max
	r.SentP90 = stat.SentBytes.P90
	r.SentP95 = stat.SentBytes.P95
	r.SentP99 = stat.SentBytes.P99

	r.ReceivedMean = stat.Bytes.Mean
	r.ReceivedMin = stat.Bytes.Min
	r.ReceivedMax = stat.Bytes.Max
	r.ReceivedP90 = stat.Bytes.P90
	r.ReceivedP95 = stat.Bytes.P95
	r.ReceivedP99 = stat.Bytes.P99

	r.ErrorCodes = maxErrors(stat.ErrorCodes, errors)
}

// RequestsStat for requests statictic (in one label)
type RequestsStat struct {
	Label string

	Stat []RequestStat

	SumStat RequestStat
}

// Sort RequestsStat by column
func (r *RequestsStat) Sort(sortCol SortColumn, sortDesc bool) {
	sort.Slice(r.Stat, func(i, j int) bool {
		switch sortCol {
		case SortSamples:
			if sortDesc {
				return r.Stat[i].Samples > r.Stat[j].Samples
			} else {
				return r.Stat[i].Samples < r.Stat[j].Samples
			}
		case SortErrors:
			if sortDesc {
				return r.Stat[i].Errors > r.Stat[j].Errors
			} else {
				return r.Stat[i].Errors < r.Stat[j].Errors
			}
		case SortResponceTimeMean:
			if sortDesc {
				return r.Stat[i].ResponceTimeMean > r.Stat[j].ResponceTimeMean
			} else {
				return r.Stat[i].ResponceTimeMean < r.Stat[j].ResponceTimeMean
			}
		case SortResponceTimeMin:
			if sortDesc {
				return r.Stat[i].ResponceTimeMin > r.Stat[j].ResponceTimeMin
			} else {
				return r.Stat[i].ResponceTimeMin < r.Stat[j].ResponceTimeMin
			}
		case SortResponceTimeMax:
			if sortDesc {
				return r.Stat[i].ResponceTimeMax > r.Stat[j].ResponceTimeMax
			} else {
				return r.Stat[i].ResponceTimeMax < r.Stat[j].ResponceTimeMax
			}
		case SortResponceTimeP90:
			if sortDesc {
				return r.Stat[i].ResponceTimeP90 > r.Stat[j].ResponceTimeP90
			} else {
				return r.Stat[i].ResponceTimeP90 < r.Stat[j].ResponceTimeP90
			}
		case SortResponceTimeP95:
			if sortDesc {
				return r.Stat[i].ResponceTimeP95 > r.Stat[j].ResponceTimeP95
			} else {
				return r.Stat[i].ResponceTimeP95 < r.Stat[j].ResponceTimeP95
			}
		case SortResponceTimeP99:
			if sortDesc {
				return r.Stat[i].ResponceTimeP99 > r.Stat[j].ResponceTimeP99
			} else {
				return r.Stat[i].ResponceTimeP99 < r.Stat[j].ResponceTimeP99
			}

		case SortSentMean:
			if sortDesc {
				return r.Stat[i].SentMean > r.Stat[j].SentMean
			} else {
				return r.Stat[i].SentMean < r.Stat[j].SentMean
			}
		case SortSentMin:
			if sortDesc {
				return r.Stat[i].SentMin > r.Stat[j].SentMin
			} else {
				return r.Stat[i].SentMin < r.Stat[j].SentMin
			}
		case SortSentMax:
			if sortDesc {
				return r.Stat[i].SentMax > r.Stat[j].SentMax
			} else {
				return r.Stat[i].SentMax < r.Stat[j].SentMax
			}
		case SortSentP90:
			if sortDesc {
				return r.Stat[i].SentP90 > r.Stat[j].SentP90
			} else {
				return r.Stat[i].SentP90 < r.Stat[j].SentP90
			}
		case SortSentP95:
			if sortDesc {
				return r.Stat[i].SentP95 > r.Stat[j].SentP95
			} else {
				return r.Stat[i].SentP95 < r.Stat[j].SentP95
			}
		case SortSentP99:
			if sortDesc {
				return r.Stat[i].SentP99 > r.Stat[j].SentP99
			} else {
				return r.Stat[i].SentP99 < r.Stat[j].SentP99
			}
		case SortReceivedMean:
			if sortDesc {
				return r.Stat[i].ReceivedMean > r.Stat[j].ReceivedMean
			} else {
				return r.Stat[i].ReceivedMean < r.Stat[j].ReceivedMean
			}
		case SortReceivedMin:
			if sortDesc {
				return r.Stat[i].ReceivedMin > r.Stat[j].ReceivedMin
			} else {
				return r.Stat[i].ReceivedMin < r.Stat[j].ReceivedMin
			}
		case SortReceivedMax:
			if sortDesc {
				return r.Stat[i].ReceivedMax > r.Stat[j].ReceivedMax
			} else {
				return r.Stat[i].ReceivedMax < r.Stat[j].ReceivedMax
			}
		case SortReceivedP90:
			if sortDesc {
				return r.Stat[i].ReceivedP90 > r.Stat[j].ReceivedP90
			} else {
				return r.Stat[i].ReceivedP90 < r.Stat[j].ReceivedP90
			}
		case SortReceivedP95:
			if sortDesc {
				return r.Stat[i].ReceivedP95 > r.Stat[j].ReceivedP95
			} else {
				return r.Stat[i].ReceivedP95 < r.Stat[j].ReceivedP95
			}
		case SortReceivedP99:
			if sortDesc {
				return r.Stat[i].ReceivedP99 > r.Stat[j].ReceivedP99
			} else {
				return r.Stat[i].ReceivedP99 < r.Stat[j].ReceivedP99
			}
		case SortRequest:
			if sortDesc {
				return strings.Compare(r.Stat[i].Request, r.Stat[j].Request) > 0
			} else {
				return strings.Compare(r.Stat[i].Request, r.Stat[j].Request) < 0
			}
		default:
			if sortDesc {
				return strings.Compare(r.Stat[i].Request, r.Stat[j].Request) > 0
			} else {
				return strings.Compare(r.Stat[i].Request, r.Stat[j].Request) < 0
			}
		}
	})
}

// LabelStat for requests statictic, grouped by labels
type LabelStat struct {
	Started int64
	Ended   int64

	Name string

	Stat []RequestsStat
}

// Init from aggstat.LabelURLAggStat
func (l *LabelStat) Init(lstat *aggstat.LabelURLAggStat, errors int) {
	l.Name = lstat.Name
	l.Started = lstat.Started
	l.Ended = lstat.Ended

	l.Stat = make([]RequestsStat, len(lstat.Stat))
	i := 0
	for label := range lstat.Stat {
		l.Stat[i].Label = label
		ustat := lstat.Stat[label]
		l.Stat[i].Stat = make([]RequestStat, len(ustat.Stat))
		j := 0
		for u := range ustat.Stat {
			l.Stat[i].Stat[j].Init(u, ustat.Stat[u], errors)
			j++
		}
		l.Stat[i].Sort(SortRequest, false)

		l.Stat[i].SumStat.Init("Summary", &ustat.SumStat, errors)

		i++
	}

	// sort labels
	sort.Slice(l.Stat, func(i, j int) bool {
		return strings.Compare(l.Stat[i].Label, l.Stat[j].Label) < 0
	})

}
