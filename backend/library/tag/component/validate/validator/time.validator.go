package validator

import (
	"fmt"
	"reflect"
	"time"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
)

type timeValidator struct{}

func (vl *timeValidator) Validate(results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
	newResults := append([]model.ValidatorResult{}, results...)

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		addNewError := func() {
			newResults = append(newResults, model.NewValidatorResult(fieldName, fieldValue, tagArgElem.ArgName, tagArgElem.ArgValue))
		}

		if tagArgElem.ArgName == enum.TagArgRequired {
			if fieldValue == nil || time.Time.IsZero(fieldValue.(time.Time)) {
				addNewError()
			}
		} else {
			return nil, fmt.Errorf("unsupported tag arg for time validator. field name: %s, arg: %s", fieldName, tagArgElem.ArgName)
		}
	}

	return newResults, nil
}

var TimeValidator = &timeValidator{}
