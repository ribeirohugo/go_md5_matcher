package main

import (
	"fmt"
	"github.com/ribeirohugo/go_md5_matcher/internal/fault"
	"github.com/ribeirohugo/go_md5_matcher/internal/matcher"
)

const (
	dataFilePath      = "csv.csv"
	dataFileColumn    = 3
	encodedFilePath   = "md5.csv"
	encodedFileColumn = 2
	delimiter         = ';'
)

func main() {
	csvMatcher := matcher.NewCsvMatcher(dataFilePath, dataFileColumn, encodedFilePath, encodedFileColumn, delimiter)

	err := csvMatcher.Open()
	if err != nil {
		fault.HandleError(err)
	}

	matched, err := csvMatcher.Match()
	if err != nil {
		fault.HandleError(err)
	}

	fmt.Println(matched)

}
