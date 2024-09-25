package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const tab = "\t|\t"

type TabWriter struct {
	Writer *tabwriter.Writer
}

func NewTabWriter() *TabWriter {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

	return &TabWriter{
		Writer: w,
	}
}

func (tw *TabWriter) WriteHeaderListCurrencies() {
	tw.Writer.Write([]byte(fmt.Sprintf("ID%sCountry%sCode%s\n", tab, tab, tab)))
}

func (tw *TabWriter) WriteListCurrencies(i int, country, code string) {
	id := i + 1
	s := fmt.Sprintf("%d%s%s%s%s%s\n", id, tab, country, tab, code, tab)

	tw.Writer.Write([]byte(s))
}

func (tw *TabWriter) WriteHeaderListRates() {
	tw.Writer.Write([]byte(fmt.Sprintf("ID%sCode%sRate%s\n", tab, tab, tab)))
}

func (tw *TabWriter) WriteListRates(i int, code string, rate float64) {
	id := i + 1
	s := fmt.Sprintf("%d%s%s%s%f%s\n", id, tab, code, tab, rate, tab)

	tw.Writer.Write([]byte(s))
}

func (tw *TabWriter) WriteHeaderConvert() {
	tw.Writer.Write([]byte(fmt.Sprintf("ID%sFrom%sTo%sAmount%sResult%s\n", tab, tab, tab, tab, tab)))
}

func (tw *TabWriter) WriteConvert(i int, from string, to string, amount float64, result float64) {
	s := fmt.Sprintf("%d%s%s%s%s%s%f%s%f%s\n", i+1, tab, from, tab, to, tab, amount, tab, result, tab)

	tw.Writer.Write([]byte(s))
}
