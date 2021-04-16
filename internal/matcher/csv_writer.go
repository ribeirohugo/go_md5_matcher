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

func (w *CsvWriter) write(row []string) error {
	return w.writer.Write(row)
}
