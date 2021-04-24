package matcher

import (
	"io/ioutil"
	"os"
	"testing"
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

	if matcher != expected {
		t.Errorf("CsvMatcher incorrect, \ngot: %v, \nwant: %v.", matcher, expected)
	}

	err := matcher.Match()

	if err != nil {
		t.Errorf("Error returned,\n got: %v,\n want: %v.", err, nil)
	}

	contentBytes, err := ioutil.ReadFile(outputFileName)
	content := string(contentBytes)
	if err != nil {
		t.Errorf("Error reading file %s", err)
	}

	if content != outputFile {
		t.Errorf("Error returned,\n got: %v,\n want: %v.", content, outputFile)
	}

	dataCsvFile.Close()
	encodedCsvFile.Close()
	outputCsvFile.Close()
}

func createTempFile(name string, content string) (*os.File, error) {

	tempFile, err := ioutil.TempFile("", name)
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
