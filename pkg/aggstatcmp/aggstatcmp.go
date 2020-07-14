package aggstatcmp

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
)

// AggDiffStat for compare aggregate stat
//easyjson:json
type AggDiffStat struct {
	Elapsed   aggstat.AggStatNode
	Connect   aggstat.AggStatNode
	Bytes     aggstat.AggStatNode
	SentBytes aggstat.AggStatNode

	Errors       float64
	SuccessCodes map[string]float64
	ErrorCodes   map[string]float64
}

func diffPcnt(s float64, o float64, max float64) float64 {
	if max == 0.0 {
		if s == 0.0 && o == 0.0 {
			return 0.0
		}
		return math.MaxFloat64
	}
	return (s - o) / max * 100.0
}

func aggStatNodeDiff(d *aggstat.AggStatNode, s *aggstat.AggStatNode, o *aggstat.AggStatNode) {
	d.Max = diffPcnt(s.Max, o.Max, o.Max)
	d.Min = diffPcnt(s.Min, o.Min, o.Max)
	d.Mean = diffPcnt(s.Mean, o.Mean, o.Max)
	d.P90 = diffPcnt(s.P90, o.P90, o.Max)
	d.P95 = diffPcnt(s.P95, o.P95, o.Max)
	d.P99 = diffPcnt(s.P99, o.P99, o.Max)
}

func diffMapPcnt(d *map[string]float64, s *map[string]uint64, sRequests uint64, o *map[string]uint64, oRequests uint64) {
	for k := range *s {
		vo, ok := (*o)[k]
		if ok {
			(*d)[k] = (float64((*s)[k])/float64(sRequests) - float64(vo)/float64(oRequests)) * 100.0
		} else {
			(*d)[k] = float64((*s)[k]) / float64(sRequests) * 100.0
		}
	}
	for k := range *o {
		_, ok := (*s)[k]
		if !ok {
			(*d)[k] = -float64((*o)[k]) / float64(oRequests) * 100.0
		}
	}
}

func sum(m map[string]uint64) (sum uint64) {
	for k := range m {
		sum += m[k]
	}
	return
}

// Init compare aggregate stat s with o
func (d *AggDiffStat) Init(s *aggstat.AggStat, o *aggstat.AggStat) {
	d.SuccessCodes = map[string]float64{}
	d.ErrorCodes = map[string]float64{}

	sErrors := sum(s.ErrorCodes)
	sRequests := sErrors + sum(s.SuccessCodes)
	oErrors := sum(o.ErrorCodes)
	oRequests := oErrors + sum(o.SuccessCodes)
	if sRequests == 0 || oRequests == 0 {
		d.Errors = math.NaN()
		d.SuccessCodes = map[string]float64{}
		d.ErrorCodes = map[string]float64{}
		d.Elapsed.NaN()
		d.Connect.NaN()
		d.Bytes.NaN()
		d.SentBytes.NaN()
	} else {
		sErrors := float64(sErrors) / float64(sRequests)
		oErrors := float64(oErrors) / float64(oRequests)
		d.Errors = (sErrors - oErrors) * 100.0

		aggStatNodeDiff(&d.Elapsed, &s.Elapsed, &o.Elapsed)
		aggStatNodeDiff(&d.Connect, &s.Connect, &o.Connect)
		aggStatNodeDiff(&d.Bytes, &s.Bytes, &o.Bytes)
		aggStatNodeDiff(&d.SentBytes, &s.SentBytes, &o.SentBytes)

		diffMapPcnt(&d.SuccessCodes, &s.SuccessCodes, sRequests, &o.SuccessCodes, oRequests)
		diffMapPcnt(&d.ErrorCodes, &s.ErrorCodes, sRequests, &o.ErrorCodes, oRequests)
	}
}

// URLAggDiffStat for compare url aggregates stat
//easyjson:json
type URLAggDiffStat struct {
	Stat map[string]*AggDiffStat

	SumStat AggDiffStat
}

// LabelURLAggDiffStat for store comparation of aggregates stat
//easyjson:json
type LabelURLAggDiffStat struct {
	Name    string
	Started int64
	Ended   int64

	CmpName    string
	CmpStarted int64
	CmpEnded   int64

	Stat map[string]*URLAggDiffStat
}

// Init compare compare aggregates stat s with o
func (d *LabelURLAggDiffStat) Init(s *aggstat.LabelURLAggStat, o *aggstat.LabelURLAggStat) {
	d.Name = s.Name
	d.Started = s.Started
	d.Ended = s.Ended

	d.CmpName = o.Name
	d.CmpStarted = o.Started
	d.CmpEnded = o.Ended
	d.Stat = map[string]*URLAggDiffStat{}
	for label := range s.Stat {
		lStat := s.Stat[label]
		loStat, loOK := o.Stat[label]
		if loOK {
			labelStat := new(URLAggDiffStat)
			labelStat.Stat = map[string]*AggDiffStat{}
			for url := range lStat.Stat {
				ou, ok := loStat.Stat[url]
				if ok {
					stat := new(AggDiffStat)
					stat.Init(lStat.Stat[url], ou)
					labelStat.Stat[url] = stat
				}
			}
			labelStat.SumStat.Init(&lStat.SumStat, &loStat.SumStat)
			d.Stat[label] = labelStat
		}
	}
}

// LabelURLAggDiffStatLoad AggStat from JSON or js file
func LabelURLAggDiffStatLoad(path string) (*LabelURLAggDiffStat, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	s := new(LabelURLAggDiffStat)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(path, ".json") {
		err = s.UnmarshalJSON(data)
	} else if strings.HasSuffix(path, ".js") {
		err = s.UnmarshalJSON(data[9:])
	} else {
		return nil, fmt.Errorf("unknown format")
	}

	return s, err
}

// Reset LabelURLAggDiffStat
func (d *LabelURLAggDiffStat) Reset() {
	d.Started = 0
	d.Ended = 0

	d.Name = ""

	d.Stat = map[string]*URLAggDiffStat{}
}
