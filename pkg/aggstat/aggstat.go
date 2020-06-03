package aggstat

import (
	"math"

	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
)

// AggStatNode for aggregates stat
//easyjson:json
type AggStatNode struct {
	Min float64
	Max float64

	Mean float64
	P90  float64
	P95  float64
	P99  float64
}

// NaN set AggStatNode fields to NaN
func (s *AggStatNode) NaN() {
	s.Min = math.NaN()
	s.Max = math.NaN()

	s.Mean = math.NaN()
	s.P90 = math.NaN()
	s.P95 = math.NaN()
	s.P99 = math.NaN()
}

// AggStat for aggregates stat
//easyjson:json
type AggStat struct {
	Started int64
	Ended   int64

	Count uint64

	Elapsed   AggStatNode
	Connect   AggStatNode
	Bytes     AggStatNode
	SentBytes AggStatNode

	Success      uint64
	SuccessCodes map[string]uint64
	ErrorCodes   map[string]uint64
}

// Init for aggregates stat
func (agg *AggStat) Init(stat *jmeterstat.JMeterStat) *AggStat {
	agg.Started = stat.Started
	agg.Ended = stat.Ended

	agg.Count = stat.Elapsed.Count()

	agg.Elapsed.Min = stat.Elapsed.Min()
	agg.Elapsed.Max = stat.Elapsed.Max()
	agg.Elapsed.Mean = stat.Elapsed.Mean()
	agg.Elapsed.P90 = stat.Elapsed.Percentile(90.0)
	agg.Elapsed.P95 = stat.Elapsed.Percentile(95.0)
	agg.Elapsed.P99 = stat.Elapsed.Percentile(99.0)

	agg.Connect.Min = stat.Connect.Min()
	agg.Connect.Max = stat.Connect.Max()
	agg.Connect.Mean = stat.Connect.Mean()
	agg.Connect.P90 = stat.Connect.Percentile(90.0)
	agg.Connect.P95 = stat.Connect.Percentile(95.0)
	agg.Connect.P99 = stat.Connect.Percentile(99.0)

	agg.Bytes.Min = stat.Bytes.Min()
	agg.Bytes.Max = stat.Bytes.Max()
	agg.Bytes.Mean = stat.Bytes.Mean()
	agg.Bytes.P90 = stat.Bytes.Percentile(90.0)
	agg.Bytes.P95 = stat.Bytes.Percentile(95.0)
	agg.Bytes.P99 = stat.Bytes.Percentile(99.0)

	agg.SentBytes.Min = stat.SentBytes.Min()
	agg.SentBytes.Max = stat.SentBytes.Max()
	agg.SentBytes.Mean = stat.SentBytes.Mean()
	agg.SentBytes.P90 = stat.SentBytes.Percentile(90.0)
	agg.SentBytes.P95 = stat.SentBytes.Percentile(95.0)
	agg.SentBytes.P99 = stat.SentBytes.Percentile(99.0)

	agg.Success = stat.Success
	agg.SuccessCodes = stat.SuccessCodes
	agg.ErrorCodes = stat.ErrorCodes

	return agg
}

// URLAggStat for aggregates stat in map[url]*AggStat
type URLAggStat struct {
	Stat map[string]*AggStat

	SumStat AggStat
}

// LabelURLAggStat for aggregates stat in map[label]*URLAggStat
//easyjson:json
type LabelURLAggStat struct {
	Started int64
	Ended   int64

	Name string

	Stat map[string]*URLAggStat
}

// Init store aggregated statistic from jmeterstat.JMeterLabelURLStat
func (s *LabelURLAggStat) Init(urlStat jmeterstat.JMeterLabelURLStat, name string) *LabelURLAggStat {
	s.Started = 0
	s.Ended = 0

	s.Name = name

	s.Stat = map[string]*URLAggStat{}
	for label, urls := range urlStat {
		s.Stat[label] = new(URLAggStat)
		s.Stat[label].Stat = map[string]*AggStat{}
		for url, stat := range urls {
			if s.Started > stat.Started || s.Started == 0 {
				s.Started = stat.Started
			}
			if s.Ended < stat.Ended {
				s.Ended = stat.Ended
			}

			s.Stat[label].Stat[url] = new(AggStat).Init(stat)
		}

		sum := jmeterstat.JMeterURLStatSum(urls, true)
		s.Stat[label].SumStat.Init(sum)
	}

	return s
}
