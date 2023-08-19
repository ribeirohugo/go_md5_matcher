package config

import (
	"os"
	"testing"
)

var configContent = `encoded_column = 0
output_name = "testFile.csv"

[data_csv]
field_delimiter = ";"
file_path = "data.csv"
match_column = 1
start_line = 1

[encoded_csv]
field_delimiter = ";"
file_path = "md5.csv"
match_column = 1
start_line = 1
`

var configTest = Config{
	DataCsv: CsvFile{
		Delimiter:   ';',
		FilePath:    "data.csv",
		MatchColumn: 1,
		StartLine:   1,
	},
	EncodedCsv: CsvFile{
		Delimiter:   ';',
		FilePath:    "md5.csv",
		MatchColumn: 1,
		StartLine:   1,
	},
	OutputName: "testFile.csv",
}

func TestConfig(t *testing.T) {

	tempFile, _ := createTempFile()

	defer os.Remove(tempFile.Name())

	cfg, _ := Load(tempFile.Name())

	if cfg != configTest {
		t.Errorf("Wrong config file output,\n got: %v,\n want: %v.", cfg, configTest)
	}

	tempFile.Close()
}

func createTempFile() (*os.File, error) {

	tempFile, err := os.CreateTemp("", "config.toml")
	if err != nil {
		return nil, err
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(configContent)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
