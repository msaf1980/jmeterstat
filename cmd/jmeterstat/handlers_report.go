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

	"github.com/msaf1980/jmeterstat/pkg/aggtable"
	"github.com/msaf1980/jmeterstat/pkg/datatables"
)

type viewStat struct {
	Report    viewReport
	ReportURL string
	Tables    []viewTable
}

// fill data for table responce for datatables.net
func fillTableReq(tableID int, stat *aggtable.LabelStat, p *datatablesParams) datatables.Responce {
	var resp datatables.Responce
	resp.Draw = p.Draw

	labelStat := stat.Stat[tableID]

	resp.RecordsTotal = len(labelStat.Stat)
	resp.RecordsFiltered = resp.RecordsTotal

	labelStat.SortRequests(aggtable.GetSortColumn(p.OrderCol), p.OrderDesc)
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
			resp.Data[j] = labelStat.Stat[i].FillRowReq()
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
					resp.Data[j] = labelStat.Stat[i].FillRowReq()
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

// fill data for table responce for datatables.net
func fillTableErr(tableID int, stat *aggtable.LabelStat, p *datatablesParams) datatables.Responce {
	var resp datatables.Responce
	resp.Draw = p.Draw

	labelStat := stat.Stat[tableID]

	resp.RecordsTotal = len(labelStat.ErrStat)

	labelStat.SortErrors(p.OrderCol, p.OrderDesc)
	start := p.Start
	length := p.Length
	if start < 0 || length <= 0 || start >= len(labelStat.ErrStat) {
		resp.Data = [][]string{}
		return resp
	} else if length > len(labelStat.ErrStat)-start {
		length = len(labelStat.ErrStat) - start
	}

	if len(p.Search) == 0 {
		resp.Data = make([][]string, length)
		j := 0
		for i := start; i < start+length; i++ {
			if len(labelStat.ErrStat[i].ErrorCodes) > 0 {
				resp.Data[j] = labelStat.ErrStat[i].FillRowErr()
				j++
			}
		}
		resp.RecordsFiltered = resp.RecordsTotal
	} else {
		resp.Data = make([][]string, length)
		j := 0
		filtered := 0
		for i := 0; i < len(labelStat.ErrStat); i++ {
			if strings.Contains(labelStat.ErrStat[i].Request, p.Search) {
				if start > 0 {
					start--
				} else if j < length {
					resp.Data[j] = labelStat.ErrStat[i].FillRowErr()
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

// return Jmeter statistics table rows
func tableRequests(stat *aggtable.LabelStat, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if n, err := strconv.Atoi(vars["id"]); err == nil {
		p, err := getDatablesParams(r)
		if err == nil {
			respData := fillTableReq(n, stat, &p)
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

// return Jmeter statistics table rows
func tableErrors(stat *aggtable.LabelStat, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if n, err := strconv.Atoi(vars["id"]); err == nil {
		p, err := getDatablesParams(r)
		if err == nil {
			respData := fillTableErr(n, stat, &p)

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

// Generate page for view Jmeter satatistics
func report(stat *aggtable.LabelStat, isCompare bool, w http.ResponseWriter, r *http.Request) {
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
			tables[i] = viewTable{Label: stat.Stat[i].Label, ID: strconv.Itoa(i), MaxErrors: 5}

			tables[i].FooterReq = stat.Stat[i].SumStat.FillRowReq()

			tables[i].Errors = make([]int, tables[i].MaxErrors)
			tables[i].FooterErr = stat.Stat[i].ErrSumStat.FillRowErr()
		}
		tmplParams := viewStat{
			Report: viewReport{
				Title:   stat.Name,
				Started: time.Unix(stat.Started/1000, 0).Format(time.RFC3339),
				Ended:   time.Unix(stat.Ended/1000, 0).Format(time.RFC3339),
			},
			Tables: tables,
		}
		if isCompare {
			tmplParams.ReportURL = "cmpreport"
		} else {
			tmplParams.ReportURL = "report"
		}
		if stats.diffStat != nil {
			tmplParams.Report.URL = true
		}
		err = tmpl.Execute(w, tmplParams)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
	} else {
		eStr := fmt.Sprintf("can't template %s for %s: %s", source, stat.Name, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(eStr))
	}
}
