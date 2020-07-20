package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime/pprof"
	"strings"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
	"github.com/msaf1980/jmeterstat/pkg/aggstatcmp"
	"github.com/msaf1980/jmeterstat/pkg/aggtable"
	"github.com/msaf1980/jmeterstat/pkg/aggtablecmp"
	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
	urltransform "github.com/msaf1980/jmeterstat/pkg/urltransform"
)

// Copy the src file to dst dir with original filename. Any existing file will be overwritten and will not
// copy file attributes.
func copyToDir(src, dstDir string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	filename := filepath.Base(src)
	out, err := os.OpenFile(path.Join(dstDir, filename), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		out.Close()
		return err
	}
	return out.Close()
}

func dumpAggStat(agg *aggstat.LabelURLAggStat, out string, htmlOut bool, jsonOut bool, template string) {
	var obytes []byte
	var ofile *os.File
	var err error

	obytes, err = agg.MarshalJSON()
	if err != nil {
		log.Fatalf("can't marshal JSON: %s", err.Error())
	}

	if jsonOut {
		ofile, err = os.OpenFile(path.Join(out, "report.json"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create JSON file in out dir: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write JSON file: %s", err.Error())
		}
	}

	if htmlOut {
		var files = []string{"report-tables.js", "report-data.js", "report.html"}
		if len(template) == 0 {
			eTemplate := "web/"
			for i := range files {
				source := eTemplate + files[i]
				data, err := Asset(source)
				if err != nil {
					log.Fatalf("can't read embedded %s: %s", source, err.Error())
				}
				ofile, err = os.OpenFile(path.Join(out, files[i]), os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					log.Fatalf("can't create %s in out dir: %s", files[i], err.Error())
				}
				_, err = ofile.Write(data)
				if err != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
				err = ofile.Close()
				if err != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
			}
		} else {
			for i := range files {
				if copyToDir(path.Join(template, files[i]), out) != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
			}
		}

		ofile, err = os.OpenFile(path.Join(out, "report.js"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create report.js in out dir: %s", err.Error())
		}
		_, err = ofile.WriteString("var data=")
		if err != nil {
			log.Fatalf("can't write report.js: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write report.js: %s", err.Error())
		}
	}
}

func readAggStat(filename string) (*aggstat.LabelURLAggStat, error) {
	agg := new(aggstat.LabelURLAggStat)

	var ibytes []byte
	var ifile *os.File
	var err error

	jsonIn := false

	/*
		obytes, err = agg.MarshalJSON()
		if err != nil {
			log.Fatalf("can't marshal JSON: %s", err.Error())
		}
	*/
	if strings.HasSuffix(filename, ".json") {
		jsonIn = true
	} else if !strings.HasSuffix(filename, ".js") {
		return nil, fmt.Errorf("unknown file type: %s", filename)
	}

	ifile, err = os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %s: %s", filename, err.Error())
	}
	ibytes, err = ioutil.ReadAll(ifile)
	if err != nil {
		log.Fatalf("can't write JSON file: %s", err.Error())
	}

	if jsonIn {
		err = agg.UnmarshalJSON(ibytes)
		if err != nil {
			return nil, fmt.Errorf("can't unmarshall file %s: %s", filename, err.Error())
		}
	} else {
		prefix := []byte("var data=")
		if bytes.HasPrefix(ibytes, prefix) {
			err = agg.UnmarshalJSON(ibytes[9:])
			if err != nil {
				return nil, fmt.Errorf("can't unmarshall file %s: %s", filename, err.Error())
			}
		}
	}

	return agg, nil
}

func readCsv(csvFilename *string, urlTransformRule urltransform.URLTransformRule) jmeterstat.JMeterLabelURLStat {
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

	return urlStat
}

func dumpDiffAggStat(diffAgg *aggstatcmp.LabelURLAggDiffStat, cmpAgg *aggstat.LabelURLAggStat, out string, htmlOut bool, jsonOut bool, template string) {
	obytes, err := diffAgg.MarshalJSON()
	if err != nil {
		log.Fatalf("can't marshal JSON: %s", err.Error())
	}

	var ofile *os.File

	if jsonOut {
		ofile, err = os.OpenFile(path.Join(out, "compare.json"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create JSON file in out dir: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write JSON file: %s", err.Error())
		}
	}

	if htmlOut {
		var files = []string{"compare-tables.js", "compare-data.js", "compare.html"}
		if len(template) == 0 {
			eTemplate := "web/"
			for i := range files {
				source := eTemplate + files[i]
				data, err := Asset(source)
				if err != nil {
					log.Fatalf("can't read embedded %s: %s", source, err.Error())
				}
				ofile, err = os.OpenFile(path.Join(out, files[i]), os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					log.Fatalf("can't create %s in out dir: %s", files[i], err.Error())
				}
				_, err = ofile.Write(data)
				if err != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
				err = ofile.Close()
				if err != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
			}
		} else {
			for i := range files {
				if copyToDir(path.Join(template, files[i]), out) != nil {
					log.Fatalf("can't write %s: %s", files[i], err.Error())
				}
			}
		}

		ofile, err = os.OpenFile(path.Join(out, "compare.js"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create compare.js in out dir: %s", err.Error())
		}
		_, err = ofile.WriteString("var data=")
		if err != nil {
			log.Fatalf("can't write compare.js: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write compare.js: %s", err.Error())
		}
	}
}

type action int

const (
	// NONE undefined
	NONE action = iota
	// CSVREPORT generate report from csv
	CSVREPORT
	// LOADREPORT load aggregated report
	LOADREPORT
	// COMPARE compare two reports
	COMPARE
	// LOADCOMPARE load diff aggregated report
	LOADCOMPARE
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	cpuProfile := fs.String("cpuprofile", "", "Write cpu profile to file")

	name := fs.String("name", "", "Test name")

	out := fs.String("out", "", "Dir for store report")

	var template string
	fs.StringVar(&template, "template", "", "Dir for html templates")

	var jsonOut bool
	var htmlOut bool
	fs.BoolVar(&jsonOut, "json", false, "save json")
	fs.BoolVar(&htmlOut, "html", false, "save html report")

	maxErrors := 5

	action := NONE

	csvFilename := fs.String("csvfile", "", "JMeter results (CSV format) (use '-' for stdin)")
	urlTransform := fs.String("urltransform", "", "Transformation rule for URL (nned for aggregate URLs stat)")

	cmpReport := fs.String("cmp", "", "jmeter report for compare")
	report := fs.String("report", "", "jmeter report")

	diffReport := fs.String("diff", "", "jmeter compare report")

	var httpListen string
	fs.StringVar(&httpListen, "http", "", "Listen embedded http server (for example :8080)")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", os.Args[0])
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n\n")

		fmt.Fprintf(os.Stderr, "Save aggregated statistic (generate html report):\n\t%s -csvfile results.csv -urltransform '{param_value.target}' -out report\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Save aggregated statistic and start embedded http server on 8080 port:\n\t%s -csvfile results.csv -json -urltransform '{param_value.target}' -out report -http ':8080'\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Load aggregated statistic and start embedded http server on 8080 port:\n\t%s -report report/report.json -http ':8080'\n", os.Args[0])

		fmt.Fprintf(os.Stderr, "\n")

		fmt.Fprintf(os.Stderr, "Compare and save aggregated statistic (generate html report):\n\t%s -report report/report.json -cmp report-cmp/report.json -out compare\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Compare and save aggregated statistic, start embedded http server on 8080 port:\n\t%s -report report/report.json -cmp report-cmp/report.json -json -out compare -http ':8080'\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Load diff aggregated statistic, start embedded http server on 8080 port:\n\t%s -diff report-compare/compare.json -http ':8080'\n", os.Args[0])

		os.Exit(1)
	}

	_ = fs.Parse(os.Args[1:])

	if !jsonOut && !htmlOut {
		jsonOut = true
	}

	if len(*csvFilename) > 0 {
		if len(*report) > 0 || len(*cmpReport) > 0 || len(*diffReport) > 0 {
			log.Fatal(fmt.Errorf("can't save and load reports in one step"))
		}
		action = CSVREPORT
	} else if len(*diffReport) > 0 {
		action = LOADCOMPARE
	} else if len(*report) > 0 {
		if len(*urlTransform) > 0 {
			log.Fatal(fmt.Errorf("unsupported option urltransform for compare or load"))
		}
		if len(*cmpReport) == 0 {
			action = LOADREPORT
		} else {
			action = COMPARE
		}
	} else {
		log.Fatal(fmt.Errorf("report not set"))
	}

	if action == NONE {
		log.Fatal(fmt.Errorf("input type not set"))
	}

	if len(*out) == 0 && httpListen == "" {
		log.Fatal(fmt.Errorf("out dir not set"))
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

	switch action {
	case CSVREPORT:
		urlTransformRule, err := urltransform.NewURLTransformRule(*urlTransform)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = os.Mkdir(*out, 0755)
		if err != nil {
			log.Fatalf("can't create out dir: %s", err.Error())
		}
		urlStat := readCsv(csvFilename, urlTransformRule)

		var agg aggstat.LabelURLAggStat
		agg.Init(urlStat, *name)
		dumpAggStat(&agg, *out, htmlOut, jsonOut, template)
		if httpListen != "" {
			var aggTable aggtable.LabelStat
			aggTable.Init(&agg, maxErrors)
			agg.Reset()
			stats.stat = &aggTable

			aggReport(httpListen)
		}

	case LOADREPORT:
		var err error
		stats.stat, err = aggtable.Load(report, maxErrors)
		if err != nil {
			log.Fatalf("can't load %s: %s", *report, err.Error())
		}
		aggReport(httpListen)

	case COMPARE:
		agg, err := readAggStat(*report)
		if err != nil {
			log.Fatal(err.Error())
		}
		cmpAgg, err := readAggStat(*cmpReport)
		if err != nil {
			log.Fatal(err.Error())
		}

		var diffAgg aggstatcmp.LabelURLAggDiffStat
		diffAgg.Init(agg, cmpAgg)
		err = os.Mkdir(*out, 0755)
		if err != nil {
			log.Fatalf("can't create out dir: %s", err.Error())
		}
		dumpDiffAggStat(&diffAgg, cmpAgg, *out, htmlOut, jsonOut, template)
		if httpListen != "" {
			stats.stat = new(aggtable.LabelStat)
			stats.stat.Init(agg, maxErrors)
			agg.Reset()

			stats.cmpStat = new(aggtable.LabelStat)
			stats.cmpStat.Init(cmpAgg, maxErrors)
			cmpAgg.Reset()

			stats.diffStat = new(aggtablecmp.LabelDiffStat)
			stats.diffStat.Init(&diffAgg, maxErrors)
			diffAgg.Reset()

			aggDiffReport(httpListen)
		}

	case LOADCOMPARE:
		var err error
		stats.diffStat, err = aggtablecmp.Load(diffReport, maxErrors)
		if err != nil {
			log.Fatalf("can't load %s: %s", *diffReport, err.Error())
		}
		stats.stat, err = aggtable.Load(report, maxErrors)
		if err != nil {
			log.Fatalf("can't load %s: %s", *report, err.Error())
		}
		stats.cmpStat, err = aggtable.Load(cmpReport, maxErrors)
		if err != nil {
			log.Fatalf("can't load %s: %s", *cmpReport, err.Error())
		}
		aggDiffReport(httpListen)
	default:
		log.Fatalf("unknown action: %d", action)
	}
}
