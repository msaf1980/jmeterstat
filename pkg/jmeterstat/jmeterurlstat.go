package jmeterstat

import (
	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
)

// JMeterURLStat map[label][url]*JMeterStat
type JMeterURLStat map[string]map[string]*JMeterStat

// JMeterURLStatAdd Add stat record to JMeterURLStat
func JMeterURLStatAdd(urlStat JMeterURLStat, url string, jmtrRecord *jmeterreader.JmtrRecord) {
	label, ok := urlStat[jmtrRecord.Label]
	if !ok {
		label = map[string]*JMeterStat{}
		urlStat[jmtrRecord.Label] = label
	}
	stat, ok := label[url]
	if !ok {
		stat = new(JMeterStat)
		stat.Init(jmtrRecord.TimeStamp)
		label[url] = stat
	}
	stat.Add(jmtrRecord.Elapsed, jmtrRecord.Connect, float64(jmtrRecord.Bytes), float64(jmtrRecord.SentBytes))
	//fmt.Printf("%s %.2f\n", url, jmtrRecord.Elapsed)
}
