package config

import (
	"fmt"
	"io"
	"os"
	"time"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
)

// CsvFile holds environment configurations about Csv files to bea read.
type CsvFile struct {
	Delimiter   rune
	FilePath    string `toml:"file_path"`
	MatchColumn int    `toml:"match_column"`
	StartLine   int    `toml:"start_line"`

	fieldDelimiter string `toml:"field_delimiter"`
}

// Config holds configurations data and methods.
type Config struct {
	DataCsv       CsvFile `toml:"data_csv"`
	EncodedCsv    CsvFile `toml:"encoded_csv"`
	EncodedColumn int     `toml:"encoded_column"`
	OutputName    string  `toml:"output_name"`
}

// Load loads configurations from a given toml file path.
func Load(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	_ = file.Close()

	config := Config{
		EncodedColumn: -1,
		DataCsv: CsvFile{
			Delimiter: ';',
		},
		EncodedCsv: CsvFile{
			Delimiter: ';',
		},
		OutputName: fmt.Sprintf("%d.csv", time.Now().Unix()),
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	if config.DataCsv.fieldDelimiter != "" {
		config.DataCsv.Delimiter, _ = utf8.DecodeRuneInString(config.DataCsv.fieldDelimiter)
	}

	if config.EncodedCsv.fieldDelimiter != "" {
		config.EncodedCsv.Delimiter, _ = utf8.DecodeRuneInString(config.EncodedCsv.fieldDelimiter)
	}

	return config, nil
}
