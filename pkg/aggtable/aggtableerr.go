package aggtable

import (
	"sort"
	"strconv"
	"strings"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
)

// ErrorCount error count
type ErrorCount struct {
	Error  string
	Errors uint64
}

// ErrorStat ErrorCodes slice
type ErrorStat struct {
	Request string

	Samples    uint64
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
func (e *ErrorStat) Init(request string, stat *aggstat.AggStat, errors int) {
	e.Request = request
	e.Samples = stat.Count

	e.ErrorCodes = maxErrors(stat.ErrorCodes, errors)
}

// FillRowErr return string table presentation
func (e *ErrorStat) FillRowErr() []string {
	data := make([]string, 2*len(e.ErrorCodes)+2)
	data[0] = e.Request
	data[1] = strconv.FormatUint(e.Samples, 10)

	for i := range e.ErrorCodes {
		if len(e.ErrorCodes[i].Error) == 0 {
			data[2+2*i] = ""
			data[3+2*i] = ""
		} else {
			data[2+2*i] = strconv.FormatUint(e.ErrorCodes[i].Errors, 10)
			data[3+2*i] = e.ErrorCodes[i].Error
		}
	}

	return data
}
