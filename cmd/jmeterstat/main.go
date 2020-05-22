package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	//"sort"

	"github.com/msaf1980/jmeterstat/pkg/aggstat"
	"github.com/msaf1980/jmeterstat/pkg/jmeterreader"
	"github.com/msaf1980/jmeterstat/pkg/jmeterstat"
	urltransform "github.com/msaf1980/jmeterstat/pkg/urltransform"
)

func searchHTMLTemplate() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can't get home dir: %s", err.Error())
	}
	_, err = os.Stat(path.Join(home, ".config", "jmeterstat", "template", "report.html"))
	if err == nil {
		return path.Join(home, ".config", "jmeterstat"), nil
	}

	if runtime.GOOS != "windows" {
		_, err = os.Stat(path.Join("/usr/local/share/jmeterstat", "template", "report.html"))
		if err == nil {
			return path.Join("/usr/local/share/jmeterstat", "template"), nil
		}
	}

	var dir string
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", fmt.Errorf("can't get executable basedir: %s", err.Error())
	}

	_, err = os.Stat(path.Join(dir, "template", "report.html"))
	if err == nil {
		return path.Join(dir, "template"), nil
	}

	_, err = os.Stat(path.Join(".", "template", "report.html"))
	if err == nil {
		return path.Join(".", "template"), nil
	}

	return "", fmt.Errorf("template dir not found")
}

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

func dump(urlStat jmeterstat.JMeterLabelURLStat, name string, out string, htmlOut bool, jsonOut bool, template string) {
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
		ofile, err = os.OpenFile(path.Join(out, "aggregate.json"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create JSON file in out dir: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write JSON file: %s", err.Error())
		}
	}

	if htmlOut {
		ofile, err = os.OpenFile(path.Join(out, "aggregate.js"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("can't create aggregate.js in out dir: %s", err.Error())
		}
		_, err = ofile.WriteString("var data=")
		if err != nil {
			log.Fatalf("can't write aggregate.js: %s", err.Error())
		}
		_, err = ofile.Write(obytes)
		if err != nil {
			log.Fatalf("can't write aggregate.js: %s", err.Error())
		}
		if copyToDir(path.Join(template, "report-tables.js"), out) != nil {
			log.Fatalf("can't write tables.js: %s", err.Error())
		}
		if copyToDir(path.Join(template, "report-data.js"), out) != nil {
			log.Fatalf("can't write report-data.js: %s", err.Error())
		}
		if copyToDir(path.Join(template, "report.html"), out) != nil {
			log.Fatalf("can't write report.html: %s", err.Error())
		}
	}
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

func main() {
	csvFilename := flag.String("file", "", "JMeter results (CSV format)")

	cpuProfile := flag.String("cpuprofile", "", "Write cpu profile to file")
	urlTransform := flag.String("urltransform", "", "Transformation rule for URL (nned for aggregate URLs stat)")

	name := flag.String("name", "", "Test name")

	out := flag.String("out", "", "Dir for store report")

	var template string
	flag.StringVar(&template, "template", "", "Dir for html templates")
	if len(template) == 0 {
		var err error
		template, err = searchHTMLTemplate()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	var jsonOut bool
	var htmlOut bool
	flag.BoolVar(&jsonOut, "json", false, "save json")
	flag.BoolVar(&htmlOut, "html", true, "save html report")

	flag.Parse()
	if !jsonOut && !htmlOut {
		log.Fatal(fmt.Errorf("output type not set"))
	}
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

	urlStat := readCsv(csvFilename, urlTransformRule)
	dump(urlStat, *name, *out, htmlOut, jsonOut, template)
}
