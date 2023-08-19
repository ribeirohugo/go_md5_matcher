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

func TestCsvMatcher_Match(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		t.Run("should match", func(t *testing.T) {
			// Load Data CSV File
			dataCsvFile, err := createTempFile(dataCsvName, dataFile)
			require.NoError(t, err)
			dataFileName := dataCsvFile.Name()

			// Load Encoded Data CSV File
			encodedCsvFile, err := createTempFile(encodedCsvName, encodedFile)
			require.NoError(t, err)
			encodedFileName := encodedCsvFile.Name()

			// Load Encoded Data CSV File
			outputCsvFile, err := createTempFile(outputCsvName, "")
			require.NoError(t, err)
			outputFileName := outputCsvFile.Name()

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

			err = matcher.Match()
			require.Nil(t, err)

			contentBytes, err := os.ReadFile(outputFileName)
			require.NoError(t, err)
			assert.Equal(t, outputFile, string(contentBytes))

			closeFile(t, dataCsvFile)
			closeFile(t, encodedCsvFile)
			closeFile(t, outputCsvFile)
		})
	})
}

func createTempFile(name string, content string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", name)
	if err != nil {
		return nil, err
	}

	_, err = tempFile.WriteString(content)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func closeFile(t *testing.T, file *os.File) {
	t.Helper()

	err := file.Close()
	require.NoError(t, err)

	err = os.Remove(file.Name())
	require.NoError(t, err)
}
