package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"

	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
	urltransform "github.com/msaf1980/jmeterstat/pkg/urltransform"
)

func readCsv(csvFilename *string, urlTransformRule urltransform.URLTransformRule) {
	csvReader, err := jmeterreader.NewJmeterCsvReader(csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	var jmtrRecord jmeterreader.JmtrRecord
	urlStat := jmeterstat.JMeterURLStat{}
	//var allThreads int

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

		jmeterstat.JMeterURLStatAdd(urlStat, url, &jmtrRecord)
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
				label, url, s.Elapsed.Count(), s.Elapsed.Max(), s.Elapsed.Min(), s.Elapsed.Mean(),
				s.Elapsed.Percentile(95), s.Elapsed.Percentile(99))
		}
	}
}

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

	readCsv(csvFilename, urlTransformRule)
}
