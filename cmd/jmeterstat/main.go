package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"

	"github.com/msaf1980/jmeterstat/pkg/jmetercsv"
	"github.com/msaf1980/jmeterstat/pkg/statcalc"
	urltransform "github.com/msaf1980/jmeterstat/pkg/urltransform"
)

func main() {
	csvFilename := flag.String("file", "", "JMeter results (CSV format)")

	cpuProfile := flag.String("cpuprofile", "", "Write cpu profile to file")
	urlTransform := flag.String("urltransform", "", "Transformation rule for URL (nned for aggregate URLs stat)")

	flag.Parse()
	if len(*csvFilename) == 0 {
		log.Fatal(fmt.Errorf("filename not set"))
	}

	if len(*cpuProfile) != 0 {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	urlTransformRule, err := urltransform.NewURLTransformRule(*urlTransform)
	if err != nil {
		log.Fatal(err.Error())
	}

	csvReader, err := jmetercsv.NewJmtrCsvReader(csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	var jmtrRecord jmetercsv.JmtrRecord
	urlStat := map[string]map[string]*statcalc.StatCalculator{}

	for {
		err = csvReader.Read(&jmtrRecord)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err.Error())
		}
		var url string
		if len(urlTransformRule) > 0 {
			url, err = urltransform.URLTransform(jmtrRecord.URL, urlTransformRule)
			if err != nil {
				log.Fatalf("%s: %s", err.Error(), jmtrRecord.URL)
			} else if len(url) == 0 {
				url = jmtrRecord.URL
			}
			//fmt.Printf("%s\n", url)
		} else {
			//fmt.Printf("%v\n", jmtrRecord)
			url = jmtrRecord.URL
		}

		label, ok := urlStat[jmtrRecord.Label]
		if !ok {
			label = map[string]*statcalc.StatCalculator{}
			urlStat[jmtrRecord.Label] = label
		}
		stat, ok := label[url]
		if !ok {
			stat = new(statcalc.StatCalculator)
			stat.Init()
			label[url] = stat
		}
		stat.AddValue(jmtrRecord.Elapsed)
		//fmt.Printf("%s %.2f\n", url, jmtrRecord.Elapsed)
	}
	labels := make([]string, 0, len(urlStat))
	for k := range urlStat {
		labels = append(labels, k)
	}
	sort.Strings(labels)
	for _, label := range labels {
		stats := urlStat[label]
		urls := make([]string, 0, len(stats))
		for k := range stats {
			urls = append(urls, k)
		}
		for _, url := range urls {
			s := stats[url]
			fmt.Printf("[%s] %s samples %d, max %.2f, min %.2f, mean %.2f, p95 %.2f, p99 %.2f\n",
				label, url, s.Count(), s.Max(), s.Min(), s.Mean(), s.Percentile(95), s.Percentile(99))
		}
	}
}
