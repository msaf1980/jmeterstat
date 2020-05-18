package jmeterstat

import (
	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
)

// JMeterURLStat map[label][url]*JMeterStat
type JMeterURLStat map[string]*JMeterStat

// JMeterLabelURLStat map[label][url]*JMeterStat
type JMeterLabelURLStat map[string]JMeterURLStat

// JMeterURLStatAdd Add stat record to JMeterURLStat
func JMeterURLStatAdd(urlStat JMeterLabelURLStat, url string, jmtrRecord *jmeterreader.JmtrRecord) {
	label, ok := urlStat[jmtrRecord.Label]
	if !ok {
		label = JMeterURLStat{}
		urlStat[jmtrRecord.Label] = label
	}
	stat, ok := label[url]
	if !ok {
		stat = new(JMeterStat)
		stat.Init()
		label[url] = stat
	}
	stat.Add(jmtrRecord.TimeStamp, jmtrRecord.Elapsed, jmtrRecord.Connect,
		float64(jmtrRecord.Bytes), float64(jmtrRecord.SentBytes),
		jmtrRecord.Success, jmtrRecord.ResponseCode)
	//fmt.Printf("%s %.2f\n", url, jmtrRecord.Elapsed)
}

// JMeterURLStatSum get summary statistic for all queries in JMeterURLStat
func JMeterURLStatSum(stat JMeterURLStat, clear bool) *JMeterStat {
	sum := new(JMeterStat)
	sum.Init()

	for _, v := range stat {
		if sum.Started > v.Started || sum.Started == 0 {
			sum.Started = v.Started
		}
		if sum.Ended < v.Ended {
			sum.Ended = v.Ended
		}
		sum.Elapsed.AddAll(&v.Elapsed)
		sum.Connect.AddAll(&v.Connect)
		sum.Bytes.AddAll(&v.Bytes)
		sum.SentBytes.AddAll(&v.SentBytes)

		sum.Success += v.Success
		for k := range v.ResponceCodes {
			sum.ResponceCodes[k] += v.ResponceCodes[k]
		}

		if clear {
			v.Init()
		}
	}

	return sum
}
