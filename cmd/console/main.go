package main

import (
	"github.com/ribeirohugo/go_md5_matcher/internal/config"
	"github.com/ribeirohugo/go_md5_matcher/internal/fault"
	"github.com/ribeirohugo/go_md5_matcher/internal/matcher"
)

const (
	configFile = "config.toml"
)

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleFatalError(err)

	dataCsv := matcher.CsvFile{
		Delimiter:   cfg.DataCsv.Delimiter,
		FilePath:    cfg.DataCsv.FilePath,
		MatchColumn: cfg.DataCsv.MatchColumn,
	}

	encodedCsv := matcher.CsvFile{
		Delimiter:   cfg.EncodedCsv.Delimiter,
		FilePath:    cfg.EncodedCsv.FilePath,
		MatchColumn: cfg.EncodedCsv.MatchColumn,
	}

	csvMatcher := matcher.NewCsvMatcher(dataCsv, encodedCsv, cfg.OutputName, cfg.EncodedColumn)

	err = csvMatcher.Match()
	if err != nil {
		fault.HandleError(err)
	}
}
