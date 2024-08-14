package transformer

import (
	"fmt"
	"reflect"
	"time"

	enum "fudjie.waizly/backend-test/library/tag/component/transform/enum"
	model "fudjie.waizly/backend-test/library/tag/component/transform/model"
)

type timeTransformer struct{}

func (trf *timeTransformer) Transform(tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) error {
	cFieldValue := fieldValue.(time.Time)

	if cFieldValue.IsZero() {
		return nil
	}

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		if tagArgElem.ArgName == enum.TagArgDate {
			// remove time
			fieldRef.Set(reflect.ValueOf(time.Date(cFieldValue.Year(), cFieldValue.Month(), cFieldValue.Day(), 0, 0, 0, 0, cFieldValue.Location())))
		} else {
			return fmt.Errorf("unsupported tag arg for time transformer. field name: %s. arg: %s", fieldName, tagArgElem.ArgName)
		}
	}

	return nil
}

var TimeTransformer = &timeTransformer{}
