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
	dataReader    *csv.Reader
	encodedReader *csv.Reader
}

func NewCsvMatcher(dataFile string, dataColumn int, encodedFile string, encodedColumn int, delimiter rune) CsvMatcher {
	return CsvMatcher{
		dataCsv: CsvFile{
			Delimiter:   delimiter,
			FilePath:    dataFile,
			MatchColumn: dataColumn,
		},

		encodedCsv: CsvFile{
			Delimiter:   delimiter,
			FilePath:    encodedFile,
			MatchColumn: encodedColumn,
		},
	}
}

func (m *CsvMatcher) Open() error {

	csvFile, err := os.Open(m.dataCsv.FilePath)
	if err != nil {
		return err
	}

	m.dataReader = csv.NewReader(csvFile)
	m.dataReader.Comma = m.dataCsv.Delimiter

	md5File, err := os.Open(m.encodedCsv.FilePath)
	if err != nil {
		return err
	}

	m.encodedReader = csv.NewReader(md5File)
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

func md5Convert(field string) string {
	toBytes := []byte(field)
	toMd5 := md5.Sum(toBytes)
	hexString := hex.EncodeToString(toMd5[:])

	return hexString
}
