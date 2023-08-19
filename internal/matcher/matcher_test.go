package matcher

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	delimiter      = ';'
	dataCsvName    = "data.csv"
	encodedColumn  = 0
	encodedCsvName = "md5.csv"
	outputCsvName  = "output.csv"
)

var dataFile = `123;123
123;123
`

var encodedFile = `202cb962ac59075b964b07152d234b70;202cb962ac59075b964b07152d234b70
202cb962ac59075b964b07152d234b70;202cb962ac59075b964b07152d234b70
`

var outputFile = `202cb962ac59075b964b07152d234b70;123;123
202cb962ac59075b964b07152d234b70;123;123
`

func TestCsvMatcher_Match(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		t.Run("should match", func(t *testing.T) {
			// Load Data CSV File
			dataCsvFile, _ := createTempFile(dataCsvName, dataFile)
			dataFileName := dataCsvFile.Name()
			defer os.Remove(dataFileName)

			// Load Encoded Data CSV File
			encodedCsvFile, _ := createTempFile(encodedCsvName, encodedFile)
			encodedFileName := encodedCsvFile.Name()
			defer os.Remove(encodedFileName)

			// Load Encoded Data CSV File
			outputCsvFile, _ := createTempFile(outputCsvName, "")
			outputFileName := outputCsvFile.Name()
			defer os.Remove(outputFileName)

			dataCsv := CsvFile{
				Delimiter:   delimiter,
				FilePath:    dataFileName,
				MatchColumn: 1,
			}

			encodedCsv := CsvFile{
				Delimiter:   delimiter,
				FilePath:    encodedFileName,
				MatchColumn: 1,
			}

			matcher := NewCsvMatcher(dataCsv, encodedCsv, outputFileName, encodedColumn)

			err := matcher.Match()
			assert.Nil(t, err)

			contentBytes, err := os.ReadFile(outputFileName)
			require.NoError(t, err)
			assert.Equal(t, outputFile, string(contentBytes))

			dataCsvFile.Close()
			encodedCsvFile.Close()
			outputCsvFile.Close()
		})
	})
}

func TestNewCsvMatcher(t *testing.T) {
	t.Run("should create a CsvMatcher successfully", func(t *testing.T) {
		const (
			dataFileName    = "data file name"
			encodedFileName = "encoded file name"
			outputFileName  = "output file name"
		)

		dataCsv := CsvFile{
			Delimiter:   delimiter,
			FilePath:    dataFileName,
			MatchColumn: 1,
		}
		encodedCsv := CsvFile{
			Delimiter:   delimiter,
			FilePath:    encodedFileName,
			MatchColumn: 1,
		}

		expected := CsvMatcher{
			dataCsv: CsvFile{
				Delimiter:   delimiter,
				FilePath:    dataFileName,
				MatchColumn: 1,
			},
			encodedCsv: CsvFile{
				Delimiter:   delimiter,
				FilePath:    encodedFileName,
				MatchColumn: 1,
			},
			writerCsv: CsvWriter{
				csv: CsvFile{
					Delimiter: delimiter,
					FilePath:  outputFileName,
				},
			},
		}

		matcher := NewCsvMatcher(dataCsv, encodedCsv, outputFileName, encodedColumn)
		assert.Equal(t, expected, matcher)
	})
}

func createTempFile(name string, content string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", name)
	if err != nil {
		return nil, err
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(content)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
