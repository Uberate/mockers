package i18n

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	CSVBlockHeader = []string{"namespace", "code"}
)

// ToCSVFile will write the i18n value to a csv file. And all the info will be writen to a file. When occur the error at
// create, the error will be return. And the value will be cover.
//
// ToCSVFile use the encoding/csv directly. About more info of the csv file, see the doc of encoding/csv.
func ToCSVFile(fileName string, instance *I18n) error {
	if instance == nil {
		return fmt.Errorf("The I18n instance is nil ")
	}

	// Build the csv info. It will spend some time.

	// get all i18n message info as an array.
	messageObjects := instance.ToMessageObjects()
	sort.SliceStable(messageObjects, func(i, j int) bool {
		if messageObjects[i].Namespace != messageObjects[j].Namespace {
			return strings.Compare(messageObjects[i].Namespace, messageObjects[j].Namespace) < 0
		}
		return strings.Compare(messageObjects[i].Code, messageObjects[j].Code) < 0
	})
	res := [][]string{
		// the res first line the header of the message.
		CSVBlockHeader,
	}

	// The resNamespaceCodeMapper will mapper the index of message of code and message. And all the ln message value is
	// in one line. The index of resNamespaceCodeMapper always start from 2.
	resNamespaceCodeMapper := map[string]map[string]int{}

	// The lnMapper is to help quick index the ln info.
	lnMapper := map[string]int{}

	// build the info of the message object.

	// the inner-index is the index of the message index in the res.
	// the index like it:
	//
	// index, namespace, code, en, zh-cn
	// 1(because the index is start from the second line), namespace_1, 00001, test, 测试
	// 2, namespace_1, 00002, (empty, in the file, it will not be written.),,
	// 3, namespace_2, 00001, ,
	// the index will be cached in the namespaceCodeInnerIndex.
	namespaceCodeInnerIndex := 1
	lnInnerIndex := 0
	for _, item := range messageObjects {
		// pre do, to check the index.
		if _, ok := resNamespaceCodeMapper[item.Namespace]; !ok {
			resNamespaceCodeMapper[item.Namespace] = map[string]int{}
		}
		if _, ok := resNamespaceCodeMapper[item.Namespace][item.Code]; !ok {
			resNamespaceCodeMapper[item.Namespace][item.Code] = namespaceCodeInnerIndex
			// insert new value of message
			namespaceCodeInnerIndex++

			res = append(res, []string{item.Namespace, item.Code})
		}
		if _, ok := lnMapper[item.Language.ToString()]; !ok {
			lnMapper[item.Language.ToString()] = lnInnerIndex + 2
			res[0] = append(res[0], item.Language.ToString())
			lnInnerIndex++
		}

		if len(res[resNamespaceCodeMapper[item.Namespace][item.Code]]) < 2+lnMapper[item.Language.ToString()] {
			// len of []string{namespace, code} is 2, but the last index is 1, so, to fill ln index - (ln() + 1)
			for i := len(res[resNamespaceCodeMapper[item.Namespace][item.Code]]) - 1; i < lnMapper[item.Language.ToString()]; i++ {
				res[resNamespaceCodeMapper[item.Namespace][item.Code]] =
					append(res[resNamespaceCodeMapper[item.Namespace][item.Code]], "")
			}
		}

		// add a message
		res[resNamespaceCodeMapper[item.Namespace][item.Code]][lnMapper[item.Language.ToString()]] = item.Message
	}

	// fill all namespace
	for i := 1; i < len(res); i++ {
		for len(res[i]) < len(res[0]) {
			res[i] = append(res[i], "")
		}
	}

	// open the csv file, and ready to write value.
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	writer := csv.NewWriter(csvFile)
	return writer.WriteAll(res)
}

// BuildFromCSVFile will build an i18n values from specify csv-file.
func BuildFromCSVFile(fileName string) (*I18n, error) {
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
