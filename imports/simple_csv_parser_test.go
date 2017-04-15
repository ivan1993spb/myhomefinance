package imports

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"testing"
)

const csvData = `
"Mon, 02 Jan 1993 00:00:00 +0300";+25482;"start";start
"Sat, 05 Nov 2016 19:39:58 +0300";-2010;"flat";flat
"Sun, 06 Nov 2016 05:39:58 +0300";-19000;"flat";flat
"Sun, 06 Nov 2016 05:39:58 +0300";-190;"fee";fee
"Mon, 07 Nov 2016 10:52:02 +0300";-30;metro;transport
"Mon, 07 Nov 2016 13:52:02 +0300";-35;"lunch";lunch
"Mon, 07 Nov 2016 22:32:13 +0300";-102;"fruit";sweet
"Mon, 07 Nov 2016 18:52:02 +0300";-30;"coffee";coffee
"Mon, 07 Nov 2016 22:33:50 +0300";-51;"chocolate";sweet
`

func DataProviderSimpleCSVParser_ReadTransaction() io.Reader {
	return bytes.NewBuffer([]byte(csvData))
}

func TestSimpleCSVParser_ReadTransaction(t *testing.T) {
	reader := &SimpleCSVParser{
		counter: 0,
		Reader:  csv.NewReader(DataProviderSimpleCSVParser_ReadTransaction()),
	}

	reader.Reader.FieldsPerRecord = 4
	reader.Reader.Comma = ';'
	//reader.Reader.Comment
	reader.Reader.LazyQuotes = true
	//reader.Reader.TrailingComma
	//reader.Reader.TrimLeadingSpace

	for {
		transaction, err := reader.ReadTransaction()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(transaction)
	}
}
