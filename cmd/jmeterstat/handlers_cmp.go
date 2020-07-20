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
	"github.com/msaf1980/jmeterstat/pkg/aggtablecmp"
	"github.com/msaf1980/jmeterstat/pkg/datatables"
)

type viewStatDiff struct {
	Report    viewReport
	CmpReport viewReport
	Tables    []viewTable
}

// fill data for table responce for datatables.net
func fillTableDiffReq(tableID int, stat *aggtablecmp.LabelDiffStat, p *datatablesParams) datatables.Responce {
	var resp datatables.Responce
	resp.Draw = p.Draw

	labelStat := stat.Stat[tableID]

	resp.RecordsTotal = len(labelStat.Stat)
	resp.RecordsFiltered = resp.RecordsTotal

	labelStat.SortRequests(aggtablecmp.GetSortColumn(p.OrderCol), p.OrderDesc)
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
func fillTableDiffErr(tableID int, stat *aggtablecmp.LabelDiffStat, p *datatablesParams) datatables.Responce {
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
func tableDiffRequests(stat *aggtablecmp.LabelDiffStat, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if n, err := strconv.Atoi(vars["id"]); err == nil {
		p, err := getDatablesParams(r)
		if err == nil {
			respData := fillTableDiffReq(n, stat, &p)
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
func tableDiffErrors(stat *aggtablecmp.LabelDiffStat, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if n, err := strconv.Atoi(vars["id"]); err == nil {
		p, err := getDatablesParams(r)
		if err == nil {
			respData := fillTableDiffErr(n, stat, &p)

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

// Generate page for view Jmeter diff satatistics
func reportDiff(diffStat *aggtablecmp.LabelDiffStat, stat *aggtable.LabelStat, cmpStat *aggtable.LabelStat, w http.ResponseWriter, r *http.Request) {
	source := "web/template/compare.html"
	t, err := Asset(source)
	if err != nil {
		eStr := fmt.Sprintf("can't read embedded %s for compare %s: %s", source, diffStat.Name, err.Error())
		log.Fatal(eStr)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(eStr))
		return
	}

	tmpl := template.New("compare")
	tmpl, err = tmpl.Parse(string(t))
	if err == nil {
		tables := make([]viewTable, len(stats.diffStat.Stat))
		for i := range stats.diffStat.Stat {
			tables[i] = viewTable{Label: stats.diffStat.Stat[i].Label, ID: strconv.Itoa(i), MaxErrors: 5}

			tables[i].FooterReq = stats.diffStat.Stat[i].SumStat.FillRowReq()

			tables[i].Errors = make([]int, tables[i].MaxErrors)
			tables[i].FooterErr = stats.diffStat.Stat[i].ErrSumStat.FillRowErr()
		}
		tmplParams := viewStatDiff{
			Report: viewReport{
				Title:   diffStat.Name,
				Started: time.Unix(diffStat.Started/1000, 0).Format(time.RFC3339),
				Ended:   time.Unix(diffStat.Ended/1000, 0).Format(time.RFC3339),
			},
			CmpReport: viewReport{
				Title:   diffStat.CmpName,
				Started: time.Unix(diffStat.CmpStarted/1000, 0).Format(time.RFC3339),
				Ended:   time.Unix(diffStat.CmpEnded/1000, 0).Format(time.RFC3339),
			},
			Tables: tables,
		}
		if stat != nil {
			tmplParams.Report.URL = true
		}
		if cmpStat != nil {
			tmplParams.CmpReport.URL = true
		}
		err = tmpl.Execute(w, tmplParams)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
	} else {
		eStr := fmt.Sprintf("can't template %s for compare %s: %s", source, stats.stat.Name, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(eStr))
	}
}
