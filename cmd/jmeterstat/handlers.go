package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/msaf1980/jmeterstat/pkg/aggtable"
	"github.com/msaf1980/jmeterstat/pkg/aggtablecmp"
)

var stats statsProcessed

type viewReport struct {
	Title   string
	Started string
	Ended   string
}

type viewTable struct {
	Label string
	ID    string

	FooterReq []string

	MaxErrors int
	Errors    []int
	FooterErr []string
}

type statsProcessed struct {
	stat     *aggtable.LabelStat
	cmpStat  *aggtable.LabelStat
	diffStat *aggtablecmp.LabelDiffStat
}

type datatablesParams struct {
	Draw        int
	Start       int
	Length      int
	SearchRegex bool
	Search      string
	OrderCol    int
	OrderDesc   bool
}

// extract datatables.net query params
func getDatablesParams(r *http.Request) (datatablesParams, error) {
	var p datatablesParams
	var err error

	params := r.URL.Query()
	if p.Draw, err = strconv.Atoi(params["draw"][0]); err != nil {
		return p, fmt.Errorf("uncorect draw")
	}
	if p.Start, err = strconv.Atoi(params["start"][0]); err != nil {
		return p, fmt.Errorf("Uncorrect start")
	}
	if p.Length, err = strconv.Atoi(params["length"][0]); err != nil {
		return p, fmt.Errorf("Uncorrect length")
	}
	if p.SearchRegex, err = strconv.ParseBool(params["search[regex]"][0]); err != nil {
		return p, fmt.Errorf("Uncorrect search[regex]")
	}
	p.Search = params["search[value]"][0]
	if p.OrderCol, err = strconv.Atoi(params["order[0][column]"][0]); err != nil {
		return p, fmt.Errorf("Uncorrect order[0][column]")
	}
	order := strings.ToLower(params["order[0][dir]"][0])
	if order == "desc" {
		p.OrderDesc = true
	} else if order == "asc" {
		p.OrderDesc = false
	} else {
		return p, fmt.Errorf("Uncorrect order[0][dir]")
	}

	return p, nil
}

func handlerAggRoot(w http.ResponseWriter, r *http.Request) {
	report(stats.stat, w, r)
}

func handlerReportTable(w http.ResponseWriter, r *http.Request) {
	tableRequests(stats.stat, w, r)
}

func handlerErrorTable(w http.ResponseWriter, r *http.Request) {
	tableErrors(stats.stat, w, r)
}

func handlerCmpReportTable(w http.ResponseWriter, r *http.Request) {
	tableRequests(stats.cmpStat, w, r)
}

func handlerCmpErrorTable(w http.ResponseWriter, r *http.Request) {
	tableErrors(stats.cmpStat, w, r)
}

func handlerAggDiffRoot(w http.ResponseWriter, r *http.Request) {
	reportDiff(stats.diffStat, w, r)
}

func handlerDiffReportTable(w http.ResponseWriter, r *http.Request) {
	tableDiffRequests(stats.diffStat, w, r)
}

func handlerDiffErrorTable(w http.ResponseWriter, r *http.Request) {
	tableDiffErrors(stats.diffStat, w, r)
}

func aggReport(listen string) {
	if stats.stat == nil {
		log.Fatal("report not loaded")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlerAggRoot)
	r.HandleFunc("/report/requests/{id}", handlerReportTable)
	r.HandleFunc("/report/errors/{id}", handlerErrorTable)

	httpServer := &http.Server{
		Addr:           listen,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        r,
	}
	fmt.Printf("Start embedded http server on %s\n", listen)
	log.Fatal(httpServer.ListenAndServe())
}

func aggDiffReport(listen string) {
	if stats.diffStat == nil {
		log.Fatal("report not loaded")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlerAggDiffRoot)

	r.HandleFunc("/report/requests/{id}", handlerReportTable)
	r.HandleFunc("/report/errors/{id}", handlerErrorTable)

	r.HandleFunc("/cmpreport/requests/{id}", handlerCmpReportTable)
	r.HandleFunc("/cmpreport/errors/{id}", handlerCmpErrorTable)

	r.HandleFunc("/diffreport/requests/{id}", handlerDiffReportTable)
	r.HandleFunc("/diffreport/errors/{id}", handlerDiffErrorTable)

	httpServer := &http.Server{
		Addr:           listen,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        r,
	}
	fmt.Printf("Start embedded http server on %s\n", listen)
	log.Fatal(httpServer.ListenAndServe())
}
