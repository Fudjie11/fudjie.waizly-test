package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
	"fudjie.waizly/backend-test/library/util"
)

type stringValidator struct{}

func (vl *stringValidator) Validate(results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
	newResults := append([]model.ValidatorResult{}, results...)

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		addNewError := func() {
			newResults = append(newResults, model.NewValidatorResult(fieldName, fieldValue, tagArgElem.ArgName, tagArgElem.ArgValue))
		}

		if tagArgElem.ArgName == enum.TagArgRequired {
			if fieldValue == nil || len(fieldValue.(string)) == 0 {
				addNewError()
			}
		} else if fieldValue != nil {
			cFieldValue := fieldValue.(string)

			if tagArgElem.ArgName == enum.TagArgNumber {
				if len(cFieldValue) == 0 {
					continue
				}

				if !util.IsValidNumber(cFieldValue) {
					addNewError()
				}

			} else if tagArgElem.ArgName == enum.TagArgMin {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("min tag must have value")
				}

				argValue, err := strconv.Atoi(tagArgElem.ArgValue)
				if err != nil {
					return nil, err
				}

				if len(cFieldValue) < argValue {
					addNewError()
				}
			} else if tagArgElem.ArgName == enum.TagArgMax {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("max tag must have value")
				}

				argValue, err := strconv.Atoi(tagArgElem.ArgValue)
				if err != nil {
					return nil, err
				}

				if len(cFieldValue) > argValue {
					addNewError()
				}
			} else {
				return nil, fmt.Errorf("unsupported tag arg for string validator. field name: %s, arg: %s", fieldName, tagArgElem.ArgName)
			}
		}
	}

	return newResults, nil
}

var StringValidator = &stringValidator{}
