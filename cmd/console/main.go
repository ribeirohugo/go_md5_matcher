package main

import (
	"github.com/ribeirohugo/go_md5_matcher/internal/fault"
	"github.com/ribeirohugo/go_md5_matcher/internal/matcher"
)

const (
	dataFilePath      = "data.csv"
	dataFileColumn    = 3
	encodedFilePath   = "md5.csv"
	encodedFileColumn = 2
	delimiter         = ';'
)

func main() {

	dataCsv := matcher.CsvFile{
		Delimiter:   delimiter,
		FilePath:    dataFilePath,
		MatchColumn: dataFileColumn,
	}

	encodedCsv := matcher.CsvFile{
		Delimiter:   delimiter,
		FilePath:    encodedFilePath,
		MatchColumn: encodedFileColumn,
	}

	csvMatcher := matcher.NewCsvMatcher(dataCsv, encodedCsv)

	err := csvMatcher.Match()
	if err != nil {
		fault.HandleError(err)
	}
}
