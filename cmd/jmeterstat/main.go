package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime/pprof"
	//"sort"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
	urltransform "github.com/msaf1980/jmeterstat/pkg/urltransform"
)

func dump(urlStat jmeterstat.JMeterLabelURLStat, out *string) {
	var agg aggstat.LabelURLAggStat
	agg.Init(urlStat)

	var obytes []byte
	var ofile *os.File
	var err error
	ofile, err = os.OpenFile(path.Join(*out, "aggregate.json"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("can't create file in out dir: %s", err.Error())
	}
	obytes, err = agg.MarshalJSON()
	if err != nil {
		log.Fatalf("can't marshal JSON: %s", err.Error())
	}
	_, err = ofile.Write(obytes)
	if err != nil {
		log.Fatalf("can't write out file: %s", err.Error())
	}

	/*
		labels := make([]string, 0, len(urlStat))
		for k := range urlStat {
			labels = append(labels, k)
		}
		sort.Strings(labels)
		for _, label := range labels {
			stats := agg.Stat[label]
			urls := make([]string, 0, len(stats))
			for k := range stats {
				urls = append(urls, k)
			}
			for _, url := range urls {
				s := stats[url]
				fmt.Printf("[%s] %s samples %d, max %.2f, min %.2f, mean %.2f, p90 %.2f, p95 %.2f, p99 %.2f\n",
					label, url, s.Count, s.Elapsed.Max, s.Elapsed.Min, s.Elapsed.Mean,
					s.Elapsed.P90, s.Elapsed.P95, s.Elapsed.P99)
			}
		} */
}

func readCsv(csvFilename *string, urlTransformRule urltransform.URLTransformRule, out *string) {
	csvReader, err := jmeterreader.NewJmeterCsvReader(csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	var jmtrRecord jmeterreader.JmtrRecord
	urlStat := jmeterstat.JMeterLabelURLStat{}
	//allThreads := 0

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

	dump(urlStat, out)
}

func main() {
	csvFilename := flag.String("file", "", "JMeter results (CSV format)")

	cpuProfile := flag.String("cpuprofile", "", "Write cpu profile to file")
	urlTransform := flag.String("urltransform", "", "Transformation rule for URL (nned for aggregate URLs stat)")

	out := flag.String("out", "", "dir for store report")

	flag.Parse()
	if len(*csvFilename) == 0 {
		log.Fatal(fmt.Errorf("filename not set"))
	}
	if len(*out) == 0 {
		log.Fatal(fmt.Errorf("out dir not set"))
	} else {
		_, err := os.Stat(*out)
		if err == nil {
			log.Fatalf("out dir already exist")
		} else if !os.IsNotExist(err) {
			log.Fatalf("can't stat out dir: %s", err.Error())
		}
		err = os.Mkdir(*out, 0755)
		if err != nil {
			log.Fatalf("can't create out dir: %s", err.Error())
		}
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

	readCsv(csvFilename, urlTransformRule, out)
}
