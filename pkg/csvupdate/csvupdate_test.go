package csvupdate

import (
	"errors"
	"testing"
)

func TestNewCSVFile(t *testing.T) {
	t.Run("expected behaviour", func(t *testing.T) {
		csv, err := newCSVFile("test_data/test.csv")
		if err != nil {
			t.Errorf("error loading file: %v", err)
		}

		if csv.Headers["name"] != 0 || csv.Headers["comment"] != 3 {
			t.Errorf("failed to map headers to int")
		}

		if len(csv.Data) != 2 {
			t.Errorf("failed to import all data rows")
		}
	})

	t.Run("with a non-existent file", func(t *testing.T) {
		_, err := newCSVFile("123")
		if err == nil {
			t.Errorf("inexistent file should trigger error")
		}
	})
}

func TestNewUpdateTool(t *testing.T) {
	tests := []struct {
		name         string
		outdated     string
		updated      string
		comparedBy   string
		fields       string
		wantErrConst error
		wantErrUpd   error
	}{
		{
			name:         "expected behaviour",
			outdated:     "test_data/test.csv",
			updated:      "test_data/update.csv",
			comparedBy:   "uuid",
			fields:       "rating,comment",
			wantErrConst: nil,
			wantErrUpd:   nil,
		},
		{
			name:         "with fields that are not in the csv headers",
			outdated:     "test_data/test.csv",
			updated:      "test_data/update.csv",
			comparedBy:   "uuids",
			fields:       "rating,comment",
			wantErrConst: errors.New("the header 'uuids' was not found in 'test_data/test.csv'"),
			wantErrUpd:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewUpdateTool(test.outdated, test.updated, test.comparedBy, test.fields)
			if err != nil && err.Error() != test.wantErrConst.Error() {
				t.Errorf("failed error test in constructor. want: %v, got: %v", test.wantErrConst, err)
			}
		})
	}
}
