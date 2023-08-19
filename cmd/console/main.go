package main

import (
	"github.com/ribeirohugo/go_md5_matcher/config"
	"github.com/ribeirohugo/go_md5_matcher/matcher"
	"log"
)

const (
	configFile = "config.toml"
)

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

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
		log.Fatalln(err)
	}
}
