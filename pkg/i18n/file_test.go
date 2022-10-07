package i18n

import (
	"fmt"
	"testing"
)

func TestToCSVFile(t *testing.T) {
	i18nInstance := NewI18nInstance(EnableI18nChange())
	i18nInstance.RegisterMessage(EN, "test", "test", "test")
	i18nInstance.RegisterMessage(ZHCN, "test", "test", "测试")
	i18nInstance.RegisterMessage(ZHCN, "test2", "test", "测试")
	i18nInstance.RegisterMessage(EN, "test", "test2", "test")

	if err := ToCSVFile("./test/test.csv", i18nInstance); err != nil {
		t.Error(err)
		return
	}
}

func TestBuildFromCSVFile(t *testing.T) {
	TestToCSVFile(t)
	if i18n, err := BuildFromCSVFile("./test/test.csv"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(i18n.ToMessageObjects())
	}

}
