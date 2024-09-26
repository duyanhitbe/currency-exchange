package helpers

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type Table struct {
	table simpletable.Table
}

func NewTable() *Table {
	return &Table{
		table: *simpletable.New(),
	}
}

func (t *Table) WriteCurrencies(rows [][]string) {
	t.table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Country"},
			{Align: simpletable.AlignCenter, Text: "Code"},
		},
	}

	for i, row := range rows {
		id := i + 1
		country := row[0]
		code := row[1]

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", id)},
			{Align: simpletable.AlignLeft, Text: country},
			{Align: simpletable.AlignLeft, Text: code},
		}

		t.table.Body.Cells = append(t.table.Body.Cells, r)
	}

	t.table.SetStyle(simpletable.StyleUnicode)
}

func (t *Table) WriteRates(rows [][]string) {
	t.table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Code"},
			{Align: simpletable.AlignCenter, Text: "Rate"},
		},
	}

	for i, row := range rows {
		id := i + 1
		code := row[0]
		rate := row[1]

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", id)},
			{Align: simpletable.AlignLeft, Text: code},
			{Align: simpletable.AlignLeft, Text: rate},
		}

		t.table.Body.Cells = append(t.table.Body.Cells, r)
	}

	t.table.SetStyle(simpletable.StyleUnicode)
}

func (t *Table) WriteConvert(rows [][]string) {
	t.table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "From"},
			{Align: simpletable.AlignCenter, Text: "To"},
			{Align: simpletable.AlignCenter, Text: "Amount"},
			{Align: simpletable.AlignCenter, Text: "Result"},
		},
	}

	for i, row := range rows {
		id := i + 1
		from := row[0]
		to := row[1]
		amount := row[2]
		result := row[3]

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", id)},
			{Align: simpletable.AlignLeft, Text: from},
			{Align: simpletable.AlignLeft, Text: to},
			{Align: simpletable.AlignLeft, Text: amount},
			{Align: simpletable.AlignLeft, Text: result},
		}

		t.table.Body.Cells = append(t.table.Body.Cells, r)
	}

	t.table.SetStyle(simpletable.StyleUnicode)
}

func (t *Table) Print() {
	t.table.Println()
}
