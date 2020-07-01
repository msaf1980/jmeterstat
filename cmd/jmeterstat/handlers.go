package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/msaf1980/jmeterstat/pkg/aggstatcmp"
	"github.com/msaf1980/jmeterstat/pkg/aggtable"
	"github.com/msaf1980/jmeterstat/pkg/datatables"
)

var stats statsProcessed

type viewTable struct {
	Label string
	ID    string
}

type viewStat struct {
	Title  string
	Tables []viewTable
}

type statsProcessed struct {
	stat     *aggtable.LabelStat
	cmpStat  *aggtable.LabelStat
	diffStat *aggstatcmp.LabelURLAggDiffStat
}

func report(stat *aggtable.LabelStat, w http.ResponseWriter, r *http.Request) {
	source := "web/template/report.html"
	t, err := Asset(source)
	if err != nil {
		eStr := fmt.Sprintf("can't read embedded %s for %s: %s", source, stat.Name, err.Error())
		log.Fatal(eStr)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(eStr))
		return
	}

	tmpl := template.New("report")
	tmpl, err = tmpl.Parse(string(t))
	if err == nil {
		tables := make([]viewTable, len(stat.Stat))
		for i := range stat.Stat {
			tables[i] = viewTable{Label: stat.Stat[i].Label, ID: strconv.Itoa(i)}
		}
		tmplParams := viewStat{Title: stats.stat.Name, Tables: tables}
		err = tmpl.Execute(w, tmplParams)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		}
	} else {
		eStr := fmt.Sprintf("can't template %s for %s: %s", source, stat.Name, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(eStr))
	}
}

type datatablesParams struct {
	Draw        int
	Start       int
	Length      int
	SearchRegex bool
	Search      string
	OrderCol    aggtable.SortColumn
	OrderDesc   bool
}

func fillTable(tableID int, stat *aggtable.LabelStat, p *datatablesParams) datatables.Responce {
	var resp datatables.Responce
	resp.Draw = p.Draw

	labelStat := stat.Stat[tableID]

	resp.RecordsTotal = len(labelStat.Stat)
	resp.RecordsFiltered = resp.RecordsTotal

	labelStat.Sort(p.OrderCol, p.OrderDesc)
	start := p.Start
	end := p.Start + p.Length
	if start >= end || end == 0 || end > len(labelStat.Stat) {
		resp.Data = [][]string{}
		return resp
	} else if len(labelStat.Stat) < end {
		end = len(labelStat.Stat)
	}

	resp.Data = make([][]string, end-start)
	j := 0
	for i := start; i < end; i++ {
		resp.Data[j] = make([]string, 21)
		resp.Data[j][0] = labelStat.Stat[i].Request
		resp.Data[j][1] = strconv.FormatUint(labelStat.Stat[i].Samples, 10)
		resp.Data[j][2] = strconv.FormatFloat(labelStat.Stat[i].Errors, 'f', 2, 64)

		resp.Data[j][3] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeMean, 'f', 2, 64)
		resp.Data[j][4] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeMin, 'f', 2, 64)
		resp.Data[j][5] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeMax, 'f', 2, 64)
		resp.Data[j][6] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeP90, 'f', 2, 64)
		resp.Data[j][7] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeP95, 'f', 2, 64)
		resp.Data[j][8] = strconv.FormatFloat(labelStat.Stat[i].ResponceTimeP99, 'f', 2, 64)

		resp.Data[j][9] = strconv.FormatFloat(labelStat.Stat[i].SentMean, 'f', 2, 64)
		resp.Data[j][10] = strconv.FormatFloat(labelStat.Stat[i].SentMin, 'f', 2, 64)
		resp.Data[j][11] = strconv.FormatFloat(labelStat.Stat[i].SentMax, 'f', 2, 64)
		resp.Data[j][12] = strconv.FormatFloat(labelStat.Stat[i].SentP90, 'f', 2, 64)
		resp.Data[j][13] = strconv.FormatFloat(labelStat.Stat[i].SentP95, 'f', 2, 64)
		resp.Data[j][14] = strconv.FormatFloat(labelStat.Stat[i].SentP99, 'f', 2, 64)

		resp.Data[j][15] = strconv.FormatFloat(labelStat.Stat[i].ReceivedMean, 'f', 2, 64)
		resp.Data[j][16] = strconv.FormatFloat(labelStat.Stat[i].ReceivedMin, 'f', 2, 64)
		resp.Data[j][17] = strconv.FormatFloat(labelStat.Stat[i].ReceivedMax, 'f', 2, 64)
		resp.Data[j][18] = strconv.FormatFloat(labelStat.Stat[i].ReceivedP90, 'f', 2, 64)
		resp.Data[j][19] = strconv.FormatFloat(labelStat.Stat[i].ReceivedP95, 'f', 2, 64)
		resp.Data[j][20] = strconv.FormatFloat(labelStat.Stat[i].ReceivedP99, 'f', 2, 64)

		j++
	}

	return resp
}

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
	if p.OrderCol, err = aggtable.GetSortColumn(params["order[0][column]"][0]); err != nil {
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

func table(stat *aggtable.LabelStat, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if n, err := strconv.Atoi(vars["id"]); err == nil {
		p, err := getDatablesParams(r)

		fmt.Printf("\nTable ID: %d, draw: %d, start %d, length %d, order: column %d desc %v, search (with regex: %v): '%s'",
			n, p.Draw, p.Start, p.Length, p.OrderCol, p.OrderDesc, p.SearchRegex, p.Search)

		if err == nil {
			respData := fillTable(n, stat, &p)

			fmt.Printf("\nReturned table ID: %d, draw: %d, count %d (%d), filtered %d",
				n, respData.Draw, respData.RecordsTotal, len(respData.Data), respData.RecordsFiltered)

			resp, err := respData.MarshalJSON()
			if err == nil {
				_, _ = w.Write(resp)
			} else {
				_, _ = w.Write([]byte(err.Error()))
			}
		} else {
			w.WriteHeader(http.StatusOK)
			respErr := datatables.ResponceError{Draw: p.Draw, Error: err.Error()}
			resp, err := respErr.MarshalJSON()
			if err != nil {
				_, _ = w.Write(resp)
			} else {
				_, _ = w.Write([]byte(err.Error()))
			}
		}
	} else {
		w.WriteHeader(http.StatusOK)
		respErr := datatables.ResponceError{Draw: 0, Error: "Uncorrect table id"}
		resp, err := respErr.MarshalJSON()
		if err != nil {
			_, _ = w.Write(resp)
		} else {
			_, _ = w.Write([]byte(err.Error()))
		}
	}
}

func handlerAggRoot(w http.ResponseWriter, r *http.Request) {
	report(stats.stat, w, r)
}

func handlerReportTable(w http.ResponseWriter, r *http.Request) {
	table(stats.stat, w, r)
}

func handlerCmpReportTable(w http.ResponseWriter, r *http.Request) {
	table(stats.cmpStat, w, r)
}

func handlerAggDiffRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Agg Diff with %s!", r.URL.Path[1:])
}

func aggReport(listen string) {
	if stats.stat == nil {
		log.Fatal("report not set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlerAggRoot)
	r.HandleFunc("/report/{id}", handlerReportTable)

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
	if stats.stat == nil {
		log.Fatal("report not set")
	}
	if stats.cmpStat == nil {
		log.Fatal("compare report not set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlerAggDiffRoot)

	r.HandleFunc("/report/{id}", handlerReportTable)
	r.HandleFunc("/cmpreport/{id}", handlerCmpReportTable)

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
