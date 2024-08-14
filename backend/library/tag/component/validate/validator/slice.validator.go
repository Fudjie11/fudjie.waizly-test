package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
)

type sliceValidator struct{}

func (vl *sliceValidator) Validate(results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
	newResults := append([]model.ValidatorResult{}, results...)

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		addNewError := func() {
			newResults = append(newResults, model.NewValidatorResult(fieldName, fieldValue, tagArgElem.ArgName, tagArgElem.ArgValue))
		}

		if tagArgElem.ArgName == enum.TagArgRequired {
			if fieldValue == nil || fieldRef.Len() == 0 {
				addNewError()
			}
		} else if fieldValue != nil {
			if tagArgElem.ArgName == enum.TagArgMin {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("min tag must have value")
				}

				argValue, err := strconv.Atoi(tagArgElem.ArgValue)
				if err != nil {
					return nil, err
				}

				if fieldRef.Len() < argValue {
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

				if fieldRef.Len() > argValue {
					addNewError()
				}
			} else {
				return nil, fmt.Errorf("unsupported tag arg for slice validator. field name: %s, arg: %s", fieldName, tagArgElem.ArgName)
			}
		}
	}

	return newResults, nil
}

var SliceValidator = &sliceValidator{}
