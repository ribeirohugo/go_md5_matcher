package matcher

import (
	"encoding/csv"
	"os"
)

type CsvWriter struct {
	csv CsvFile

	file   *os.File
	writer *csv.Writer
}

func (w *CsvWriter) open() error {
	var err error

	w.file, err = os.Create(w.csv.FilePath)
	if err != nil {
		return err
	}

	w.writer = csv.NewWriter(w.file)
	w.writer.Comma = w.csv.Delimiter

	return nil
}

func (w *CsvWriter) close() error {
	w.writer.Flush()
	return w.file.Close()
}

func (w *CsvWriter) write(row []string, value string) error {
	newRow := row

	if value != "" {
		newRow = prepend(newRow, value)
	}

	return w.writer.Write(newRow)
}

func prepend(row []string, value string) []string {
	row = append(row, "")
	copy(row[1:], row)
	row[0] = value
	return row
}
