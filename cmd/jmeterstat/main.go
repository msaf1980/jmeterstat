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

func dumpAggStat(urlStat jmeterstat.JMeterLabelURLStat, name string, out string, htmlOut bool, jsonOut bool, template string) {
	var agg aggstat.LabelURLAggStat
	agg.Init(urlStat, name)

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
			eTemplate := "template/"
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

func dumpDiffAggStat(agg *aggstat.LabelURLAggStat, cmpAgg *aggstat.LabelURLAggStat, out string, htmlOut bool, jsonOut bool, template string) {
	var diffAgg aggstatcmp.LabelURLAggDiffStat
	diffAgg.DiffPcnt(agg, cmpAgg)

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
			eTemplate := "template/"
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
	// COMPARE compare two reports
	COMPARE
)

func main() {
	cpuProfile := flag.String("cpuprofile", "", "Write cpu profile to file")

	name := flag.String("name", "", "Test name")

	out := flag.String("out", "", "Dir for store report")

	var template string
	flag.StringVar(&template, "template", "", "Dir for html templates")
	/* 	if len(template) == 0 {
		var err error
		template, err = searchHTMLTemplate()
		if err != nil {
			log.Fatal(err.Error())
		}
	} */

	var jsonOut bool
	var htmlOut bool
	flag.BoolVar(&jsonOut, "json", false, "save json")
	flag.BoolVar(&htmlOut, "html", true, "save html report")

	var action action

	csvFilename := flag.String("csvfile", "", "JMeter results (CSV format) (use '-' for stdin)")
	urlTransform := flag.String("urltransform", "", "Transformation rule for URL (nned for aggregate URLs stat)")

	cmpReport := flag.String("cmp", "", "jmeter report for compare")
	report := flag.String("report", "", "jmeter report")

	flag.Parse()
	if !jsonOut && !htmlOut {
		log.Fatal(fmt.Errorf("output type not set"))
	}

	if len(*csvFilename) > 0 {
		if len(*report) > 0 || len(*cmpReport) > 0 {
			log.Fatal(fmt.Errorf("set compare report in generate step"))
		}
		action = CSVREPORT
	}
	if len(*report) > 0 {
		if action != NONE {
			log.Fatal(fmt.Errorf("can't run several actions in one step"))
		}
		if len(*urlTransform) > 0 {
			log.Fatal(fmt.Errorf("unsupported option urltransform for compare"))
		}
		if len(*cmpReport) == 0 {
			log.Fatal(fmt.Errorf("compare report not set"))
		}
		action = COMPARE
	}

	if action == NONE {
		log.Fatal(fmt.Errorf("input type not set"))
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

	switch action {
	case CSVREPORT:
		urlTransformRule, err := urltransform.NewURLTransformRule(*urlTransform)
		if err != nil {
			log.Fatal(err.Error())
		}
		urlStat := readCsv(csvFilename, urlTransformRule)
		dumpAggStat(urlStat, *name, *out, htmlOut, jsonOut, template)

	case COMPARE:
		agg, err := readAggStat(*report)
		if err != nil {
			log.Fatal(err.Error())
		}
		cmpAgg, err := readAggStat(*cmpReport)
		if err != nil {
			log.Fatal(err.Error())
		}
		dumpDiffAggStat(agg, cmpAgg, *out, htmlOut, jsonOut, template)

	default:
		log.Fatalf("unknown action: %d", action)
	}
}
