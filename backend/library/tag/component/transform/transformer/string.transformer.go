package transformer

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	enum "fudjie.waizly/backend-test/library/tag/component/transform/enum"
	model "fudjie.waizly/backend-test/library/tag/component/transform/model"
)

type stringTransformer struct{}

func (trf *stringTransformer) Transform(tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) error {
	cFieldValue := fieldValue.(string)

	if len(cFieldValue) == 0 {
		return nil
	}

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		if tagArgElem.ArgName == enum.TagArgUpper {
			cFieldValue = strings.ToUpper(cFieldValue)
			fieldRef.Set(reflect.ValueOf(cFieldValue))
		} else if tagArgElem.ArgName == enum.TagArgLower {
			cFieldValue = strings.ToLower(cFieldValue)
			fieldRef.Set(reflect.ValueOf(cFieldValue))
		} else if tagArgElem.ArgName == enum.TagArgTitle {
			cFieldValue = cases.Title(language.English, cases.NoLower).String(cFieldValue)
			fieldRef.Set(reflect.ValueOf(cFieldValue))
		} else if tagArgElem.ArgName == enum.TagArgTrim {
			cFieldValue = strings.TrimSpace(cFieldValue)
			fieldRef.Set(reflect.ValueOf(cFieldValue))
		} else if tagArgElem.ArgName == enum.TagArgNoSpace {
			cFieldValue = strings.ReplaceAll(cFieldValue, " ", "")
			fieldRef.Set(reflect.ValueOf(cFieldValue))
		} else {
			return fmt.Errorf("unsupported tag arg for string transformer. field name: %s. arg: %s", fieldName, tagArgElem.ArgName)
		}
	}

	return nil
}

var StringTransformer = &stringTransformer{}
