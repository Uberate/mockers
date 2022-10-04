package i18n

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	CSVBlockHeader = []string{"namespace", "code"}
)

// BuildFromCSVDir will build an i18n values from specify csv-file.
func BuildFromCSVDir(fileName string) (*I18n, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)
	csvReader := csv.NewReader(csvFile)

	csvValue, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	// get ln header
	if len(csvValue) == 0 {
		return nil, fmt.Errorf("The csv file is empty: %s ", fileName)
	}
	if len(csvValue[0]) < 3 {
		return nil, fmt.Errorf("The csv file must has one ln at least and the block value ")
	}

	lns := csvValue[0][2:]
	var languageKeys []LanguageKey
	for _, languageName := range lns {
		languageKeys = append(languageKeys, GetLanguageKey(languageName))
	}

	// generator i18n value

	res := &I18n{}
	// build messages
	for _, item := range csvValue[1:] {
		if len(item) < 3 {
			return nil, fmt.Errorf("The message must has namespace and code value ")
		}
		namespace := item[0]
		code := item[1]
		for index, ln := range languageKeys {
			m := MessageObject{}
			if len(item) < 2+index {
				// has no message
				continue
			}
			if len(item[2+index]) == 0 {
				// has no message
				continue
			}
			m.Message = item[2+index]
			m.Namespace = namespace
			m.Code = code
			m.Language = ln
			res.RegisterMessageObject(m)
		}
	}

	return res, nil
}
