package matcher

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

type CsvFile struct {
	Delimiter   rune
	FilePath    string
	MatchColumn int
}

type CsvMatcher struct {
	dataCsv       CsvFile
	encodedCsv    CsvFile
	dataFile      *os.File
	encodedFile   *os.File
	dataReader    *csv.Reader
	encodedReader *csv.Reader
}

func NewCsvMatcher(dataFile string, dataColumn int, dataDelimiter rune, encodedFile string, encodedColumn int, encodedDelimiter rune) CsvMatcher {
	return CsvMatcher{
		dataCsv: CsvFile{
			Delimiter:   dataDelimiter,
			FilePath:    dataFile,
			MatchColumn: dataColumn,
		},

		encodedCsv: CsvFile{
			Delimiter:   encodedDelimiter,
			FilePath:    encodedFile,
			MatchColumn: encodedColumn,
		},
	}
}

func (m *CsvMatcher) Open() error {
	var err error

	m.dataFile, err = os.Open(m.dataCsv.FilePath)
	if err != nil {
		return err
	}

	m.dataReader = csv.NewReader(m.dataFile)
	m.dataReader.Comma = m.dataCsv.Delimiter

	m.encodedFile, err = os.Open(m.encodedCsv.FilePath)
	if err != nil {
		return err
	}

	m.encodedReader = csv.NewReader(m.encodedFile)
	m.encodedReader.Comma = m.encodedCsv.Delimiter

	return nil
}

func (m *CsvMatcher) Match() ([][]string, error) {
	var result [][]string

	dataLines, err := m.dataReader.ReadAll()
	if err != nil {
		return result, err
	}

	encodedLines, err := m.encodedReader.ReadAll()
	if err != nil {
		return result, err
	}

	for _, encodedLine := range encodedLines {
		field := encodedLine[m.encodedCsv.MatchColumn]

		if field != "" {
			for _, dataLine := range dataLines {
				field = strings.ToLower(field)
				dataEncoded := md5Convert(dataLine[m.dataCsv.MatchColumn])

				if field == dataEncoded {
					log.Println(field, " = ", dataEncoded)
					log.Println(dataLine)
					result = append(result, encodedLine)
					break
				}
			}
		}
	}
	return result, nil
}

func (m *CsvMatcher) Close() (error, error) {
	dataError := m.dataFile.Close()
	encodedError := m.encodedFile.Close()
	return dataError, encodedError
}

func md5Convert(field string) string {
	toBytes := []byte(field)
	toMd5 := md5.Sum(toBytes)
	hexString := hex.EncodeToString(toMd5[:])

	return hexString
}
