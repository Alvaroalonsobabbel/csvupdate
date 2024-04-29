package csvupdate

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type UpdateTool struct {
	OutdatedCSV, UpdatedCSV *CSVFile
	CompareBy               string
	UpdateFields            []string
}

type CSVFile struct {
	Header   []string
	Headers  map[string]int
	Data     [][]string
	FileName string
}

// NewUpdateTool takes both the outdated and updated csv, the relevant
// fields and returns an instance of UpdateTool with both CSVs parsed.
func NewUpdateTool(outdated, updated, compareby, fields string) (*UpdateTool, error) {
	updateTool := &UpdateTool{
		CompareBy:    compareby,
		UpdateFields: strings.Split(fields, ","),
	}
	var err error
	updateTool.OutdatedCSV, err = newCSVFile(outdated)
	if err != nil {
		return nil, err
	}

	updateTool.UpdatedCSV, err = newCSVFile(updated)
	if err != nil {
		return nil, err
	}

	if err := updateTool.checkFields(); err != nil {
		return nil, err
	}

	return updateTool, nil
}

func (u *UpdateTool) checkFields() error {
	receivedHeaders := u.UpdateFields
	receivedHeaders = append(receivedHeaders, u.CompareBy)
	csv := []*CSVFile{u.OutdatedCSV, u.UpdatedCSV}

	for _, c := range csv {
		for _, r := range receivedHeaders {
			_, ok := c.Headers[r]
			if !ok {
				return fmt.Errorf("the header '%s' was not found in '%s'", r, c.FileName)
			}
		}
	}

	return nil
}

func newCSVFile(csvFile string) (*CSVFile, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var csv = &CSVFile{
		Headers:  make(map[string]int),
		Data:     [][]string{},
		FileName: csvFile,
	}

	record, err := reader.Read()
	if err != nil {
		return nil, err
	}
	csv.Header = record
	for index, header := range record {
		csv.Headers[header] = index
	}

	csv.Data, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return csv, nil
}

func (u *UpdateTool) UpdateCSV() error {
	for _, outdated := range u.OutdatedCSV.Data {
		for _, updated := range u.UpdatedCSV.Data {
			if outdated[u.OutdatedCSV.Headers[u.CompareBy]] == updated[u.UpdatedCSV.Headers[u.CompareBy]] {
				for _, field := range u.UpdateFields {
					outdated[u.OutdatedCSV.Headers[field]] = updated[u.UpdatedCSV.Headers[field]]
				}
			}
		}
	}

	if err := u.OutdatedCSV.WriteCSV(); err != nil {
		return err
	}
	return nil
}

func (c *CSVFile) WriteCSV() error {
	const outputFileName = "out.csv"

	file, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(c.Header); err != nil {
		return err
	}

	for _, record := range c.Data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
