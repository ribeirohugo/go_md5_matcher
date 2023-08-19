package matcher

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

const outputDelimiter = ';'

// CsvFile holds Csv file data and reading configurations.
type CsvFile struct {
	Delimiter   rune
	FilePath    string
	MatchColumn int
	StartLine   int

	file   *os.File
	reader *csv.Reader
}

// CsvMatcher wraps all data about Csv files, column and writer.
type CsvMatcher struct {
	dataCsv       CsvFile
	encodedCsv    CsvFile
	encodedColumn int
	writerCsv     CsvWriter
}

// NewCsvMatcher instantiates a new CsvMatcher struct.
// Insert a CsvFile for the non encoded data.
// Insert a CsvFile for Csv MD5 encoded data to check matches.
func NewCsvMatcher(dataCsv CsvFile, encodedCsv CsvFile, outputPath string, encodedColumn int) CsvMatcher {
	return CsvMatcher{
		dataCsv:       dataCsv,
		encodedCsv:    encodedCsv,
		encodedColumn: encodedColumn,
		writerCsv: CsvWriter{
			csv: CsvFile{
				Delimiter: outputDelimiter,
				FilePath:  outputPath,
			},
		},
	}
}

func (m *CsvMatcher) open() error {
	var err error

	m.dataCsv.file, err = os.Open(m.dataCsv.FilePath)
	if err != nil {
		return err
	}

	m.dataCsv.reader = csv.NewReader(m.dataCsv.file)
	m.dataCsv.reader.Comma = m.dataCsv.Delimiter

	m.encodedCsv.file, err = os.Open(m.encodedCsv.FilePath)
	if err != nil {
		return err
	}

	m.encodedCsv.reader = csv.NewReader(m.encodedCsv.file)
	m.encodedCsv.reader.Comma = m.encodedCsv.Delimiter

	err = m.writerCsv.open()
	if err != nil {
		return err
	}

	return nil
}

// Match is run to find field matches between Csv data file and Csv MD5 encoded data file
// It will generate a new Csv file with the data file rows that matched with the encoded ones.
func (m *CsvMatcher) Match() error {
	err := m.open()
	if err != nil {
		return err
	}

	dataLines, err := m.dataCsv.reader.ReadAll()
	if err != nil {
		return err
	}

	encodedLines, err := m.encodedCsv.reader.ReadAll()
	if err != nil {
		return err
	}

	dataLength := len(dataLines)
	encodedLength := len(encodedLines)

	for i := m.encodedCsv.StartLine; i < encodedLength; i++ {
		field := encodedLines[i][m.encodedCsv.MatchColumn]

		if field != "" {
			for j := m.dataCsv.StartLine; j < dataLength; j++ {
				field = strings.ToLower(field)
				dataEncoded := md5Convert(dataLines[j][m.dataCsv.MatchColumn])

				if field == dataEncoded {
					logger := fmt.Sprintf("line %d: %s = %s", i, field, dataEncoded)
					log.Println(logger)

					encodedColumn := ""
					if m.encodedColumn >= 0 {
						encodedColumn = encodedLines[i][m.encodedColumn]
					}

					err = m.writerCsv.write(dataLines[j], encodedColumn)
					if err != nil {
						return err
					}
					m.writerCsv.writer.Flush()
					break
				}
			}
		}
	}

	err = m.dataCsv.file.Close()
	if err != nil {
		return err
	}

	err = m.encodedCsv.file.Close()
	if err != nil {
		return err
	}

	err = m.writerCsv.close()
	if err != nil {
		return err
	}

	return nil
}

func md5Convert(field string) string {
	toBytes := []byte(field)
	toMd5 := md5.Sum(toBytes)
	hexString := hex.EncodeToString(toMd5[:])

	return hexString
}
