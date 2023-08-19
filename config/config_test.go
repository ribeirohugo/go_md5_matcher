package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestLoad(t *testing.T) {
	expectedConfig := Config{
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

	t.Run("with all fields", func(t *testing.T) {
		tempFile := createTempFile(t, configContent)

		cfg, err := Load(tempFile.Name())
		require.NoError(t, err)
		assert.Equal(t, expectedConfig, cfg)

		closeFile(t, tempFile)
	})
}

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "config.toml")
	require.NoError(t, err)

	_, err = tempFile.WriteString(fileContent)
	require.NoError(t, err)

	return tempFile
}

func closeFile(t *testing.T, file *os.File) {
	t.Helper()

	err := file.Close()
	require.NoError(t, err)

	err = os.Remove(file.Name())
	require.NoError(t, err)
}
