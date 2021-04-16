package main

import (
	"log"

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
	csvMatcher := matcher.NewCsvMatcher(dataFilePath, dataFileColumn, delimiter, encodedFilePath, encodedFileColumn, delimiter)

	err := csvMatcher.Open()
	if err != nil {
		fault.HandleError(err)
	}

	matched, err := csvMatcher.Match()
	if err != nil {
		fault.HandleError(err)
	}

	log.Println(matched)

	errData, errEncoded := csvMatcher.Close()
	fault.HandleFatalError(errData)
	fault.HandleFatalError(errEncoded)
}
