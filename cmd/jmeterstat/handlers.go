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
	Label  string
	ID     string
	Footer []string
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

// Generate page for view Jmeter satatistics
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
			tables[i].Footer = fillRow(&stat.Stat[i].SumStat)
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

func fillRow(r *aggtable.RequestStat) []string {
	data := make([]string, 21)
	data[0] = r.Request
	data[1] = strconv.FormatUint(r.Samples, 10)
	data[2] = strconv.FormatFloat(r.Errors, 'f', 2, 64)

	data[3] = strconv.FormatFloat(r.ResponceTimeMean, 'f', 2, 64)
	data[4] = strconv.FormatFloat(r.ResponceTimeMin, 'f', 2, 64)
	data[5] = strconv.FormatFloat(r.ResponceTimeMax, 'f', 2, 64)
	data[6] = strconv.FormatFloat(r.ResponceTimeP90, 'f', 2, 64)
	data[7] = strconv.FormatFloat(r.ResponceTimeP95, 'f', 2, 64)
	data[8] = strconv.FormatFloat(r.ResponceTimeP99, 'f', 2, 64)

	data[9] = strconv.FormatFloat(r.SentMean, 'f', 2, 64)
	data[10] = strconv.FormatFloat(r.SentMin, 'f', 2, 64)
	data[11] = strconv.FormatFloat(r.SentMax, 'f', 2, 64)
	data[12] = strconv.FormatFloat(r.SentP90, 'f', 2, 64)
	data[13] = strconv.FormatFloat(r.SentP95, 'f', 2, 64)
	data[14] = strconv.FormatFloat(r.SentP99, 'f', 2, 64)

	data[15] = strconv.FormatFloat(r.ReceivedMean, 'f', 2, 64)
	data[16] = strconv.FormatFloat(r.ReceivedMin, 'f', 2, 64)
	data[17] = strconv.FormatFloat(r.ReceivedMax, 'f', 2, 64)
	data[18] = strconv.FormatFloat(r.ReceivedP90, 'f', 2, 64)
	data[19] = strconv.FormatFloat(r.ReceivedP95, 'f', 2, 64)
	data[20] = strconv.FormatFloat(r.ReceivedP99, 'f', 2, 64)

	return data
}

// fill data for table responce for datatables.net
func fillTable(tableID int, stat *aggtable.LabelStat, p *datatablesParams) datatables.Responce {
	var resp datatables.Responce
	resp.Draw = p.Draw

	labelStat := stat.Stat[tableID]

	resp.RecordsTotal = len(labelStat.Stat)
	resp.RecordsFiltered = resp.RecordsTotal

	labelStat.Sort(p.OrderCol, p.OrderDesc)
	start := p.Start
	length := p.Length
	if start < 0 || length <= 0 || start >= len(labelStat.Stat) {
		resp.Data = [][]string{}
		return resp
	} else if length > len(labelStat.Stat)-start {
		length = len(labelStat.Stat) - start
	}

	if len(p.Search) == 0 {
		resp.Data = make([][]string, length)
		j := 0
		for i := start; i < start+length; i++ {
			resp.Data[j] = fillRow(&labelStat.Stat[i])
			j++
		}
	} else {
		resp.Data = make([][]string, length)
		j := 0
		filtered := 0
		for i := 0; i < len(labelStat.Stat); i++ {
			if strings.Contains(labelStat.Stat[i].Request, p.Search) {
				if start > 0 {
					start--
				} else if j < length {
					resp.Data[j] = fillRow(&labelStat.Stat[i])
					j++
				}
				filtered++
			}
		}
		resp.RecordsFiltered = filtered
		if j < length {
			// not full, resize needed
			data := resp.Data
			resp.Data = make([][]string, j)
			for i := 0; i < j; i++ {
				resp.Data[i] = data[i]
			}
		}
	}

	return resp
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

// return Jmeter statistics table rows
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
