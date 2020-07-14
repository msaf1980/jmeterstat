package aggtablecmp

import (
	"sort"
	"strconv"
	"strings"

	"github.com/msaf1980/jmeterstat/pkg/aggstatcmp"
)

// ErrorCount error count
type ErrorCount struct {
	Error  string
	Errors float64
}

// ErrorDiffStat ErrorCodes slice
type ErrorDiffStat struct {
	Request string

	Samples    float64
	ErrorCodes []ErrorCount
}

func maxErrors(errCount map[string]float64, maxCount int) []ErrorCount {
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
func (e *ErrorDiffStat) Init(request string, stat *aggstatcmp.AggDiffStat, samples float64, errors int) {
	e.Request = request
	e.Samples = samples

	e.ErrorCodes = maxErrors(stat.ErrorCodes, errors)
}

// FillRowErr return string table presentation
func (e *ErrorDiffStat) FillRowErr() []string {
	data := make([]string, 2*len(e.ErrorCodes)+2)
	data[0] = e.Request
	data[1] = strconv.FormatFloat(e.Samples, 'f', 2, 64)

	for i := range e.ErrorCodes {
		if len(e.ErrorCodes[i].Error) == 0 {
			data[2+2*i] = ""
			data[3+2*i] = ""
		} else {
			data[2+2*i] = strconv.FormatFloat(e.ErrorCodes[i].Errors, 'f', 2, 64)
			data[3+2*i] = e.ErrorCodes[i].Error
		}
	}

	return data
}
