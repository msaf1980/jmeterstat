package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	//"github.com/msaf1980/jmetersstat/pkg/jmetercsv"
)

func main() {
	csvFilename := flag.String("file", "", "JMeter results (CSV format)")

	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to file")

	flag.Parse()
	if len(*csvFilename) == 0 {
		log.Fatal(fmt.Errorf("filename not set"))
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.ReuseRecord = true
	line, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", line)
		break
	}
}
