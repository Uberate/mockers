package files

import "encoding/json"

// JsonFileOperator can read or write a json file. The base is used json.Marshal() and json.Unmarshal().
type JsonFileOperator struct {
}

func (j JsonFileOperator) Read(path string, obj any) error {
	fileValue, err := ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileValue, obj)
}

func (j JsonFileOperator) Write(path string, obj any) error {
	fileValue, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return WriteFile(path, fileValue)
}
