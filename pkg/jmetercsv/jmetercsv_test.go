package jmetercsv

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"io"
	"testing"
)

func TestCorrecJtmtrCsv(t *testing.T) {
	csvFilename := "test/correct.csv"
	c, err := NewJmtrCsvReader(&csvFilename)
	if err != nil {
		t.Error(err)
	}
	var r JmtrRecord
	err = c.Read(&r)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int64(1589307723495), r.TimeStamp, "timeStamp mismatch")
	assert.Equal(t, float64(73.0), r.Elapsed, "elapsed mismatch")
	assert.Equal(t, "HTTP Request (1 Hour)", r.Label, "label mismatch")
	assert.Equal(t, "200", r.ResponseCode, "responseCode mismatch")
	assert.Equal(t, true, r.Success, "sucess mismatch")
	assert.Equal(t, int64(75), r.Bytes, "bytes mismatch")
	assert.Equal(t, int64(249), r.SentBytes, "sentBytes mismatch")
	assert.Equal(t, int64(10), r.GrpThreads, "grpThreads mismatch")
	assert.Equal(t, int64(16), r.AllThreads, "allThreads mismatch")
	assert.Equal(t, "http://127.0.0.1/render/", r.URL, "URL mismatch")
	assert.Equal(t, float64(70.0), r.Latency, "Latency mismatch")
	assert.Equal(t, float64(0.0), r.IdleTime, "IdleTime mismatch")
	assert.Equal(t, float64(34.0), r.Connect, "Connect mismatch")
}

func TestIncompleteHeaderJtmtrCsv(t *testing.T) {
	csvFilename := "test/incomplte_header.csv"
	_, err := NewJmtrCsvReader(&csvFilename)
	if err == nil {
		err = fmt.Errorf("imcomplete header verify error")
	}
	if err.Error() != "missing field: Connect" {
		t.Error(err)
	}
}

func TestWrongFieldJtmtrCsv(t *testing.T) {
	csvFilename := "test/wrong_field.csv"
	_, err := NewJmtrCsvReader(&csvFilename)
	if err == nil {
		err = fmt.Errorf("imcomplete header verify error")
	}
	if err.Error() != "unknown field: SleepTime" {
		t.Error(err)
	}
}

func TestOverFieldJtmtrCsv(t *testing.T) {
	csvFilename := "test/over_field.csv"
	_, err := NewJmtrCsvReader(&csvFilename)
	if err == nil {
		err = fmt.Errorf("imcomplete header verify error")
	}
	if err.Error() != "mismatch fields count in header" {
		t.Error(err)
	}
}

func TestNoFieldRecordJtmtrCsv(t *testing.T) {
	csvFilename := "test/no_field_record.csv"
	c, err := NewJmtrCsvReader(&csvFilename)
	if err != nil {
		t.Error(err)
	}
	var r JmtrRecord
	err = c.Read(&r)
	if err == nil {
		err = fmt.Errorf("record fields count verify error")
	}
	if err.Error() != "record on line 2: wrong number of fields" {
		t.Error(err)
	}
}

func BenchmarkJmtrCsvReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		csvFilename := "test/results.csv"
		c, err := NewJmtrCsvReader(&csvFilename)
		if err != nil {
			b.Fatal(err)
		}
		var r JmtrRecord
		for {
			err = c.Read(&r)
			if err == io.EOF {
				break
			} else if err != nil {
				b.Fatal(err)
			}
		}
	}
}
