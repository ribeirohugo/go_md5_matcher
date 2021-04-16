package matcher

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/ribeirohugo/go_md5_matcher/internal/fault"
	"log"
	"os"
	"strings"
)

const outputDelimiter = ';'

type CsvFile struct {
	Delimiter   rune
	FilePath    string
	MatchColumn int

	file   *os.File
	reader *csv.Reader
}

type CsvMatcher struct {
	dataCsv    CsvFile
	encodedCsv CsvFile
	writerCsv  CsvWriter
}

// Instantiate a new CsvMatcher struct
// Insert a CsvFile for the non encoded data
// Insert a CsvFile for Csv MD5 encoded data to check matches
func NewCsvMatcher(dataCsv CsvFile, encodedCsv CsvFile) CsvMatcher {
	timer := fault.RealTimer{}
	outputPath := fmt.Sprintf("%d.csv", timer.Now())

	return CsvMatcher{
		dataCsv:    dataCsv,
		encodedCsv: encodedCsv,
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

// Run Match to find field matches between Csv data file and Csv MD5 encoded data file
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

	for i, encodedLine := range encodedLines {
		field := encodedLine[m.encodedCsv.MatchColumn]

		if field != "" {
			for _, dataLine := range dataLines {
				field = strings.ToLower(field)
				dataEncoded := md5Convert(dataLine[m.dataCsv.MatchColumn])

				if field == dataEncoded {
					logger := fmt.Sprintf("line %d: %s = %s", i, field, dataEncoded)
					log.Println(logger)

					err = m.writerCsv.write(dataLine)
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

	return err
}

func md5Convert(field string) string {
	toBytes := []byte(field)
	toMd5 := md5.Sum(toBytes)
	hexString := hex.EncodeToString(toMd5[:])

	return hexString
}
